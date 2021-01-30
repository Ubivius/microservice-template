package data

import (
	"encoding/json"
	"io"
)

func (product *Product) FromProductJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

// ToProductJSON serializes products into a json String
func (products *Products) ToProductJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

// ToJSON serializes interface into a json String
func (products *Products) ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}
