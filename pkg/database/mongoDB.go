package database

import (
	"context"
	"os"
	"time"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		log.Error(err, "Failed to connect to database. Shutting down service")
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err, "Failed to ping database. Shutting down service")
		os.Exit(1)
	}

	log.Info("Connection to MongoDB established")

	collection := client.Database("test").Collection("products")

	// Assign client and collection to the MongoProducts struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoProducts) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err, "Error while disconnecting from database")
	}
}

func (mp *MongoProducts) GetProducts() data.Products {
	// products will hold the array of Products
	var products data.Products

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err, "Error getting products from database")
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
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
	cursor.Close(context.TODO())

	return products
}

func (mp *MongoProducts) GetProductByID(id string) (*data.Product, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Holds search result
	var result data.Product

	// Find a single matching item from the database
	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	// Parse result into the returned product
	return &result, err
}

func (mp *MongoProducts) UpdateProduct(product *data.Product) error {
	// Set updated timestamp in product
	product.UpdatedOn = time.Now().UTC().String()

	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: product.ID}}

	// Update sets the matched products in the database to product
	update := bson.M{"$set": product}

	// Update a single item in the database with the values in update that match the filter
	_, err := mp.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error(err, "Error updating product.")
	}

	return err
}

func (mp *MongoProducts) AddProduct(product *data.Product) error {
	product.ID = uuid.NewString()
	// Adding time information to new product
	product.CreatedOn = time.Now().UTC().String()
	product.UpdatedOn = time.Now().UTC().String()

	// Inserting the new product into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return err
	}
	log.Info("Inserting product", "Inserted ID", insertResult.InsertedID)

	return nil
}

func (mp *MongoProducts) DeleteProduct(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Error(err, "Error deleting product")
	}
	log.Info("Deleted documents in products collection", "delete_count", result.DeletedCount)

	return nil
}
