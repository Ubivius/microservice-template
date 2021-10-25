package database

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func integrationTestSetup(t *testing.T) {
	t.Log("Test setup")

	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	newLogger := zap.New(zap.UseFlagOptions(&opts), zap.WriteTo(os.Stdout))
	logf.SetLogger(newLogger.WithName("log"))

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

	products := mp.GetProducts(context.Background())
	return products[0].ID
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
