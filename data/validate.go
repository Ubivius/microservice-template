package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// Validate a product with json validation and customer SKU validator
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
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
