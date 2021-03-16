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
	// Check the logs to make sure that the value is inserted (you can see the inserted item id in the logs)
	mp.CloseDB()
}

// Update product in database, should be included in integration test instead
func TestMongoDBUpdateProduct(t *testing.T) {
	product := &data.Product{
		ID:          0,
		Name:        "testName",
		Description: "testDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	mp := NewMongoProducts()
	err := mp.UpdateProduct(product)
	if err != nil {
		t.Fail()
	}
	// Check the logs to make sure that the value is inserted (you can see the inserted item id in the logs)
	mp.CloseDB()
}
