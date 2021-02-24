package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ErrorProductNotFound : Product specific errors
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// Database connection
func GetDBConnection() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return nil
}

// DB client
func DBConnection() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:test@cluster0.sbzzm.mongodb.net/products?retryWrites=true&w=majority")) // Connecting to url
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) // If no error, connection settings
	err = client.Connect(ctx)                                           // Connect
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx) // Defer disconnect to end of function

	// Ping db to make sure we can access it
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// List databases in atlas
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	return nil
}

// All of these functions will become database calls in the future
// GETTING PRODUCTS

// GetProducts returns the list of products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product with the given id
func GetProductByID(id int) (*Product, error) {
	index := findIndexByProductID(id)
	if index == -1 {
		return nil, ErrorProductNotFound
	}
	return productList[index], nil
}

// UPDATING PRODUCTS

// UpdateProduct updates the product specified in received JSON
func UpdateProduct(product *Product) error {
	index := findIndexByProductID(product.ID)
	if index == -1 {
		return ErrorProductNotFound
	}
	productList[index] = product
	return nil
}

// AddProduct creates a new product
func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

// DeleteProduct deletes the product with the given id
func DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	productList = append(productList[:index], productList[index+1])

	return nil
}

// Returns the index of a product in the database
// Returns -1 when no product is found
func findIndexByProductID(id int) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// productList is a hard coded list of products for this
// example data source. Should be replaced by database connection
var productList = []*Product{
	{
		ID:          1,
		Name:        "Sword",
		Description: "A basic steel sword",
		Price:       250,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Boots",
		Description: "Simple leather boots",
		Price:       100,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
