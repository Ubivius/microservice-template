package database

import (
	"context"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
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
	err := mp.AddProduct(context.Background(), product)
	if err != nil {
		t.Errorf("Failed to add product to database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateProductIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	product := &data.Product{
		ID:          uuid.NewString(),
		Name:        "testName",
		Description: "testDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	mp := NewMongoProducts()
	err := mp.UpdateProduct(context.Background(), product)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetProductsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts()
	products := mp.GetProducts(context.Background())
	if products == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts()
	_, err := mp.GetProductByID(context.Background(), "e2382ea2-b5fa-4506-aa9d-d338aa52af44")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
