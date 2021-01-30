package data

import (
	"fmt"
	"time"
)

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

func UpdateProduct(id int, p *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[position] = p
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func AddProduct(product *Product) {
	product.ID = getNextId()
	productList = append(productList, product)
}

func getNextId() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////// Fake database ///////////////////////////////////
///// DB connection setup and docker file will be done in sprint 8 /////////
///////////////////////////////////////////////////////////////////////////
// Returns the list of products
func GetProducts() Products {
	return productList
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
