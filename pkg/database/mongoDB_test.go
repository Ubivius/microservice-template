package database

import (
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/Ubivius/shared-resource-manager/resources"
	"github.com/google/uuid"
)

func newResourceManager() resources.ResourceManager {
	return resources.NewMockResources()
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts(newResourceManager())
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

	mp := NewMongoProducts(newResourceManager())
	err := mp.AddProduct(product)
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

	mp := NewMongoProducts(newResourceManager())
	err := mp.UpdateProduct(product)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetProductsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts(newResourceManager())
	products := mp.GetProducts()
	if products == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts(newResourceManager())
	_, err := mp.GetProductByID("e2382ea2-b5fa-4506-aa9d-d338aa52af44")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
