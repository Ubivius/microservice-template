package data

import (
	"fmt"
	"time"
)

// Product specific errors
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product.
// Formatting done with json tags to the right. "-" : don't include when encoding to json
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// GETTING PRODUCTS

// Returns the list of products
// Will be a database call in the future
func GetProducts() Products {
	return productList
}

// UPDATING PRODUCTS

// need to remove id int from parameters when product handler is updated
// Will be a database call in the future
func UpdateProduct(id int, p *Product) error {
	index := findIndexByProductID(p.ID)
	if index == -1 {
		return ErrorProductNotFound
	}
	productList[index] = p
	return nil
}

// ADD A PRODUCT
// Will be a database call in the future
func AddProduct(product *Product) {
	product.ID = getNextId()
	productList = append(productList, product)
}

// DELETING A PRODUCT
// Will be a database call in the future
func DeleteProduct(id int) error {
	index := findIndexByProductID(id)
	if index == -1 {
		return ErrorProductNotFound
	}

	// This should not work, probably needs ':' after index+1. To test
	productList = append(productList[:index], productList[index+1])

	return nil
}

// Returns the index of a product in the database
// Returns -1 when no product is found
// Will be a database call in the future
func findIndexByProductID(id int) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}
	return -1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////

// Finds the maximum index of our fake database and adds 1
func getNextId() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// productList is a hard coded list of products for this
// example data source. Should be replaced by database connection
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
