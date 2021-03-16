package database

import (
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
)

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddProductIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

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

func TestMongoDBUpdateProductIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

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
