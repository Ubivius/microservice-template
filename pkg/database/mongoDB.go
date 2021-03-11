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

	log.Println("Connected to MongoDB!")
	return nil
}

func (mp *MongoProducts) CloseDB() error {
	return nil
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
	product.ID = getNextID()
	productList = append(productList, product)
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
