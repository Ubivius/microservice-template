package database

import (
	"github.com/Ubivius/microservice-template/pkg/data"
)

type MongoProducts struct {
}

func NewMongoProducts() *MongoProducts {
	return &MongoProducts{}
}

func (mp *MongoProducts) Connect() error {
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
