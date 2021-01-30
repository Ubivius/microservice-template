package data

import (
	"encoding/json"
	"io"
)

func (product *Product) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

// ToProductJSON serializes products into a json String
func (products *Products) ToProductJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}
