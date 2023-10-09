package validation

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func InitializeValidator() {
	Validate = validator.New()
}
