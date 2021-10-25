package database

import (
	"context"
	"os"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "admin")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "pass")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "27888")
	}
	if os.Getenv("DB_HOSTNAME") == "" {
		os.Setenv("DB_HOSTNAME", "localhost")
	}

	err := deleteAllProductsFromMongoDB()
	if err != nil {
		t.Errorf("Failed to delete existing items from database during setup")
	}
}

func addProductAndGetId(t *testing.T) string {
	t.Log("Adding product")
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

	t.Log("Fetching new product ID")
	products := mp.GetProducts(context.Background())
	mp.CloseDB()
	return products[len(products)-1].ID
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

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
	integrationTestSetup(t)

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

	products := mp.GetProducts(context.Background())
	if len(products) != 1 {
		t.Errorf("Added product missing from database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateNonExistantProductIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	product := &data.Product{
		ID:          uuid.NewString(),
		Name:        "testName",
		Description: "testDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	mp := NewMongoProducts()
	err := mp.UpdateProduct(context.Background(), product)
	if err == nil || err.Error() != "no matches found" {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBUpdateExistingProductIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	product := &data.Product{
		ID:          addProductAndGetId(t),
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

func TestMongoDBGetProductsFromEmptyDatabaseIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoProducts()
	products := mp.GetProducts(context.Background())
	if products != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetProductsFromNonEmptyDatabaseIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	var ids [5]string
	for i := 0; i < 5; i++ {
		ids[i] = addProductAndGetId(t)
	}

	mp := NewMongoProducts()
	products := mp.GetProducts(context.Background())

	if products == nil || len(products) != 5 {
		t.Fail()
	}

	for i := 0; i < 5; i++ {
		if ids[i] != products[i].ID {
			t.Fail()
		}
	}

	mp.CloseDB()
}

func TestMongoDBGetNonExistantProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	mp := NewMongoProducts()
	_, err := mp.GetProductByID(context.Background(), uuid.NewString())
	t.Log(err.Error())
	if err == nil || err.Error() != "mongo: no documents in result" {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetExistingProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	integrationTestSetup(t)

	id := addProductAndGetId(t)

	mp := NewMongoProducts()
	product, err := mp.GetProductByID(context.Background(), id)
	if err != nil {
		t.Fail()
	}

	if id != product.ID {
		t.Fail()
	}

	mp.CloseDB()
}
