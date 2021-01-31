package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// The current setup works well with a single struct to validate
// The struct to validate should be passed as an interface in the future and the errors should be handled as individual error strings
// For further information see :
// Validator library : https://github.com/go-playground/validator
// Nic Jackson episode : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/product-api/data/validation.go

// ValidateProduct a product with json validation and customer SKU validator
func (product *Product) ValidateProduct() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(product)
}

// Custom SKU validator
func validateSKU(fieldLevel validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fieldLevel.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
