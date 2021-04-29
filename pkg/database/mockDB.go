package database

import (
	"time"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
)

type MockProducts struct {
}

func NewMockProducts() ProductDB {
	log.Info("Connecting to mock database")
	return &MockProducts{}
}

func (mp *MockProducts) Connect() error {
	return nil
}

func (mp *MockProducts) PingDB() error {
	return nil
}

func (mp *MockProducts) CloseDB() {
	log.Info("Mocked DB connection closed")
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
	product.ID = uuid.NewString()
	productList = append(productList, product)
	return nil
}

func (mp *MockProducts) DeleteProduct(id string) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return data.ErrorProductNotFound
	}

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

var productList = []*data.Product{
	{
		ID:          "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:        "Sword",
		Description: "A basic steel sword",
		Price:       250,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          "e2382ea2-b5fa-4506-aa9d-d338aa52af44",
		Name:        "Boots",
		Description: "Simple leather boots",
		Price:       100,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
