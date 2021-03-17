package database

import (
	"log"
	"time"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
)

type MockProducts struct {
}

func NewMockProducts() ProductDB {
	return &MockProducts{}
}

func (mp *MockProducts) Connect() error {
	return nil
}

func (mp *MockProducts) CloseDB() {
	log.Println("Mocked DB connection closed")
}

func (mp *MockProducts) GetProducts() data.Products {
	return productList
}

func (mp *MockProducts) GetProductByID(id string) (*data.Product, error) {
	index := findIndexByProductID(id)
	if index == -1 {
		return nil, data.ErrorProductNotFound
	}
	return productList[index], nil
}

func (mp *MockProducts) UpdateProduct(product *data.Product) error {
	index := findIndexByProductID(product.ID)
	if index == -1 {
		return data.ErrorProductNotFound
	}
	productList[index] = product
	return nil
}

func (mp *MockProducts) AddProduct(product *data.Product) error {
	product.ID = getNextID()
	productList = append(productList, product)
	return nil
}

func (mp *MockProducts) DeleteProduct(id string) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return data.ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	productList = append(productList[:index], productList[index+1:]...)

	return nil
}

// Returns the index of a product in the database
// Returns -1 when no product is found
func findIndexByProductID(id string) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////// Mocked database ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our mocked database and adds 1
func getNextID() string {
	return uuid.NewString()
}

var productList = []*data.Product{
	{
		ID:          uuid.NewString(),
		Name:        "Sword",
		Description: "A basic steel sword",
		Price:       250,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          uuid.NewString(),
		Name:        "Boots",
		Description: "Simple leather boots",
		Price:       100,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
