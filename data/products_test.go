package data

import "testing"

func TestChecksValidation(t *testing.T) {
	product := &Product{}

	err := product.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
