package database

import (
	"log"
	"os"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
)

// TODO sprint 11: need setup step to set database to desired state before tests.
// TODO sprint 11: complete integration tests once setup task is completed

func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts(NewTestLogger())
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

	mp := NewMongoProducts(NewTestLogger())
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

	mp := NewMongoProducts(NewTestLogger())
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

	mp := NewMongoProducts(NewTestLogger())
	products := mp.GetProducts()
	log.Println("Id of first product: ", products[0].ID)
	mp.CloseDB()
	t.Fail()
}

func TestMongoDBGetProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoProducts(NewTestLogger())
	product, err := mp.GetProductByID("e2382ea2-b5fa-4506-aa9d-d338aa52af44")
	if err != nil {
		t.Fail()
	}
	log.Println("Id of first product: ", product.ID)
	mp.CloseDB()
	t.Fail()
}
