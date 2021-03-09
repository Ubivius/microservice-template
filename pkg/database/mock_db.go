package database

import (
	"github.com/Ubivius/microservice-template/pkg/data"
)

type MockProducts struct {
}

func NewMockProducts() *MockProducts {
	return &MockProducts{}
}

func (mp *MockProducts) Connect() error {
	return nil
}

func (mp *MockProducts) CloseDB() error {
	return nil
}

func (mp *MockProducts) GetProducts() data.Products {
	return productList
}

func (mp *MockProducts) GetProductByID(id int) (*data.Product, error) {
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

func (mp *MockProducts) AddProduct(product *data.Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

func (mp *MockProducts) DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return data.ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	productList = append(productList[:index], productList[index+1])

	return nil
}

// Returns the index of a product in the database
// Returns -1 when no product is found
// func findIndexByProductID(id int) int {
// 	for index, product := range productList {
// 		if product.ID == id {
// 			return index
// 		}
// 	}
// 	return -1
// }

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
// func getNextID() int {
// 	lastProduct := productList[len(productList)-1]
// 	return lastProduct.ID + 1
// }

// var productList = []*data.Product{
// 	{
// 		ID:          1,
// 		Name:        "Sword",
// 		Description: "A basic steel sword",
// 		Price:       250,
// 		SKU:         "abc323",
// 		CreatedOn:   time.Now().UTC().String(),
// 		UpdatedOn:   time.Now().UTC().String(),
// 	},
// 	{
// 		ID:          2,
// 		Name:        "Boots",
// 		Description: "Simple leather boots",
// 		Price:       100,
// 		SKU:         "fjd34",
// 		CreatedOn:   time.Now().UTC().String(),
// 		UpdatedOn:   time.Now().UTC().String(),
// 	},
// }
