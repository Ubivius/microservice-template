package database

import "github.com/Ubivius/microservice-template/pkg/data"

type MongoProducts struct {
}

func NewMongoProducts() *MongoProducts {
	return &MongoProducts{}
}

func (mp *MongoProducts) GetProducts() data.Products {
	return nil
}

func (mp *MongoProducts) GetProductByID() (*data.Product, error) {
	return &data.Product{}, nil
}
