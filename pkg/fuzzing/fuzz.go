package fuzzing

import "encoding/json"

type Product struct {
	ID          string  `json:"id" bson:"_id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

func Fuzz(data []byte) int {
	productInterface := Product{}
	product := &Product{
		Name:  "Malcolm",
		Price: 2.00,
		SKU:   "abs-abs-abscd",
	}
	prod, err := json.Marshal(product)
	if err != nil {
		if prod != nil {
			panic("product != nil on error")
		}
		return 0
	}
	err = json.Unmarshal(data, &productInterface)

	if err != nil {
		panic(err)
	}
	return 1
}
