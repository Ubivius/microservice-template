package data

import (
	"encoding/json"
	"io"
)

// FromProductJSON deserializes product from JSON string
func (product *Product) FromProductJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

// ToProductJSON serializes products into a json String
func (products *Products) ToProductJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

// FromProductJSON deserializes the interface from JSON string
func FromJSON(i interface{}, reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(i)
}

// ToJSON serializes interface into a json String
func ToJSON(i interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(i)
}
