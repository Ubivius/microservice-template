package database

import (
	"github.com/Ubivius/microservice-template/pkg/data"
)

// The interface that any kind of database must implement
type ProductDB interface {
	GetProducts() data.Products
	GetProductByID(id string) (*data.Product, error)
	UpdateProduct(product *data.Product) error
	AddProduct(product *data.Product) error
	DeleteProduct(id string) error
	Connect() error
	PingDB() error
	CloseDB()
}
