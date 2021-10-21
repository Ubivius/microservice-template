package database

import (
	"context"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// The interface that any kind of database must implement
type ProductDB interface {
	GetProducts(ctx context.Context) data.Products
	GetProductByID(ctx context.Context, id string) (*data.Product, error)
	UpdateProduct(ctx context.Context, product *data.Product) error
	AddProduct(ctx context.Context, product *data.Product) error
	DeleteProduct(ctx context.Context, id string) error
	Connect() error
	PingDB() error
	CloseDB()
}
