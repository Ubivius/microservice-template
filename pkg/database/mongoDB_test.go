package database

import (
	"flag"
	"os"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/google/uuid"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// TODO sprint 11: need setup step to set database to desired state before tests.
// TODO sprint 11: complete integration tests once setup task is completed
func setup() {
	opts := zap.Options{}
	opts.BindFlags(flag.CommandLine)
	newLogger := zap.New(zap.UseFlagOptions(&opts), zap.WriteTo(os.Stdout))
	logf.SetLogger(newLogger.WithName("zap"))
}

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	setup()

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
	setup()

	product := &data.Product{
		Name:        "testName",
		Description: "testDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	mp := NewMongoProducts()
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
	setup()

	product := &data.Product{
		ID:          uuid.NewString(),
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
	mp.CloseDB()
}

func TestMongoDBGetProductsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	setup()

	mp := NewMongoProducts()
	products := mp.GetProducts()
	if products == nil {
		t.Fail()
	}

	mp.CloseDB()
	t.Fail()
}

func TestMongoDBGetProductByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}
	setup()

	mp := NewMongoProducts()
	product, err := mp.GetProductByID("e2382ea2-b5fa-4506-aa9d-d338aa52af44")
	if err != nil || product == nil {
		t.Fail()
	}

	mp.CloseDB()
	t.Fail()
}
