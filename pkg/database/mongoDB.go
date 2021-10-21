package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// ErrorEnvVar : Environment variable error
var ErrorEnvVar = fmt.Errorf("missing environment variable")

type MongoProducts struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoProducts() ProductDB {
	mp := &MongoProducts{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		log.Error(err, "MongoDB setup failed")
		os.Exit(1)
	}
	return mp
}

func (mp *MongoProducts) Connect() error {
	uri := mongodbURI()

	// Setting client options
	opts := options.Client()
	clientOptions := opts.ApplyURI(uri)
	opts.Monitor = otelmongo.NewMonitor()

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("ubivius").Collection("products")

	// Assign client and collection to the MongoProducts struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoProducts) PingDB() error {
	return mp.client.Ping(context.Background(), nil)
}

func (mp *MongoProducts) CloseDB() {
	err := mp.client.Disconnect(context.Background())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoProducts) GetProducts(ctx context.Context) data.Products {
	// products will hold the array of Products
	var products data.Products

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err, "Error getting products from database")
	}

	// Iterating through cursor
	for cursor.Next(ctx) {
		var result data.Product
		err := cursor.Decode(&result)
		if err != nil {
			log.Error(err, "Error decoding product from database")
		}
		products = append(products, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err, "Error in cursor after iteration")
	}

	// Close the cursor once finished
	cursor.Close(ctx)

	return products
}

func (mp *MongoProducts) GetProductByID(ctx context.Context, id string) (*data.Product, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Product

	// Find a single matching item from the database
	err := mp.collection.FindOne(ctx, filter).Decode(&result)

	// Parse result into the returned product
	return &result, err
}

func (mp *MongoProducts) UpdateProduct(ctx context.Context, product *data.Product) error {
	// Set updated timestamp in product
	product.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: product.ID}}

	// Update sets the matched products in the database to product
	update := bson.M{"$set": product}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error(err, "Error updating product.")
	}

	return err
}

func (mp *MongoProducts) AddProduct(ctx context.Context, product *data.Product) error {
	product.ID = uuid.NewString()
	// Adding time information to new product
	product.CreatedOn = time.Now().UTC().String()
	product.UpdatedOn = time.Now().UTC().String()

	// Inserting the new product into the database
	insertResult, err := mp.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	log.Info("Inserting product", "Inserted ID", insertResult.InsertedID)

	return nil
}

func (mp *MongoProducts) DeleteProduct(ctx context.Context, id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err, "Error deleting product")
	}
	log.Info("Deleted documents in products collection", "delete_count", result.DeletedCount)

	return nil
}

func mongodbURI() string {
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	if hostname == "" || port == "" || username == "" || password == "" {
		log.Error(ErrorEnvVar, "Some environment variables are not available for the DB connection. DB_HOSTNAME, DB_PORT, DB_USERNAME, DB_PASSWORD")
		os.Exit(1)
	}

	return "mongodb://" + username + ":" + password + "@" + hostname + ":" + port + "/?authSource=admin"
}
