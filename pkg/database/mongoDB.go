package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Ubivius/microservice-template/pkg/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProducts struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// Should pass logger here
// NewMongoProducts(logger) to log err with log.fatal
func NewMongoProducts() ProductDB {
	mp := &MongoProducts{}
	err := mp.Connect()
	// If connect fails, kill the program
	if err != nil {
		os.Exit(1)
	}
	return mp
}

func (mp *MongoProducts) Connect() error {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:test@cluster0.sbzzm.mongodb.net/products?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil || client == nil {
		os.Exit(1)
	}

	// Ping DB
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
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
		log.Println(err)
	} else {
		log.Println("Connection to MongoDB closed.")
	}
}

func (mp *MongoProducts) GetProducts() data.Products {
	// results will hold the array of Products
	var results data.Products

	// Find returns a cursor that must be iterated through
	cursor, err := mp.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through cursor
	for cursor.Next(context.TODO()) {
		var elem data.Product
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cursor.Close(context.TODO())

	return results
}

func (mp *MongoProducts) GetProductByID(id int) (*data.Product, error) {
	// MongoDB search filter
	filter := bson.D{{Key: "id", Value: 0}}

	// Holds search result
	var result data.Product

	err := mp.collection.FindOne(context.TODO(), filter).Decode(&result)

	return &result, err
}

func (mp *MongoProducts) UpdateProduct(product *data.Product) error {
	index := findIndexByProductID(product.ID)
	if index == -1 {
		return data.ErrorProductNotFound
	}
	productList[index] = product
	return nil
}

func (mp *MongoProducts) AddProduct(product *data.Product) {
	// Adding time information to new product
	product.CreatedOn = time.Now().UTC().String()
	product.UpdatedOn = time.Now().UTC().String()

	// Inserting the new product into the database
	insertResult, err := mp.collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (mp *MongoProducts) DeleteProduct(id int) error {
	result, err := mp.collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %v documents in the products collection\n", result.DeletedCount)
	return nil
}
