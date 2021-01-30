package data

import (
	"encoding/json"
	"io"
)

func (product *Product) FromJson(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

func (products *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}
