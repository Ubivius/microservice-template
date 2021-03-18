package database

import (
	"context"
	"log"
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
	logger     *log.Logger
}

func NewMongoProducts(l *log.Logger) ProductDB {
	mp := &MongoProducts{logger: l}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		mp.logger.Fatal(err)
	}
	return mp
}

func (mp *MongoProducts) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:pass@localhost:27888/?authSource=admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		mp.logger.Fatalln("Failed to connect to database. Shutting down service")
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		mp.logger.Fatal(err)
	}

	log.Println("Connection to MongoDB established")

	collection := client.Database("test").Collection("products")

	// Assign client and collection to the MongoProducts struct
	mp.collection = collection
	mp.client = client
	return nil
}

func (mp *MongoProducts) CloseDB() {
	err := mp.client.Disconnect(context.TODO())
	if err != nil {
		mp.logger.Println(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}

func (mp *MongoProducts) GetProducts() data.Products {
	// products will hold the array of Products
	var products data.Products

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var result data.Product
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, &result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
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
		log.Println(err)
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

	log.Println("Inserting a document: ", insertResult.InsertedID)
	return nil
}

func (mp *MongoProducts) DeleteProduct(id string) error {
	// MongoDB search filter
	filter := bson.D{{Key: "_id", Value: id}}

	// Delete a single item matching the filter
	result, err := mp.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents in the products collection\n", result.DeletedCount)
	return nil
}
