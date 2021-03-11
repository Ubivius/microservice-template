package database

import (
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// Should be an integration test
// Temporary test
func TestMongoDBConnectionAndShutdown(t *testing.T) {
	mp := NewMongoProducts()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

// Add product to database, should be an integration test
func TestMongoDBAddProduct(t *testing.T) {
	product := &data.Product{
		Name:        "testName",
		Description: "testDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	mp := NewMongoProducts()
	mp.AddProduct(product)
	mp.CloseDB()
}
