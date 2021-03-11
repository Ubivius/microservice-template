package database

import (
	"context"
	"log"
	"os"

	"github.com/Ubivius/microservice-template/pkg/data"
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
	return productList
}

func (mp *MongoProducts) GetProductByID(id int) (*data.Product, error) {
	index := findIndexByProductID(id)
	if index == -1 {
		return nil, data.ErrorProductNotFound
	}
	return productList[index], nil
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
	insertResult, err := mp.collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted a single document: ", insertResult.InsertedID)
}

func (mp *MongoProducts) DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return data.ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	productList = append(productList[:index], productList[index+1])

	return nil
}
