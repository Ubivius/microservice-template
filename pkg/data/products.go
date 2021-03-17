package data

import (
	"fmt"
)

// ErrorProductNotFound : Product specific errors
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Product struct {
	ID          string  `json:"id" bson:"_id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product
