package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validate *validator.Validate

func InitializeValidator() {
	Validate = validator.New()
}

func ValidateStruct(model interface{}) error {
	err := Validate.Struct(model)
	if err != nil {
		var messageErrors strings.Builder
		messageErrors.WriteString("Validation error ")
		for _, err := range err.(validator.ValidationErrors) {
			messageErrors.WriteString(err.Field() + ": " + err.Tag() + ", ")
		}
		return errors.New(messageErrors.String())
	}
	return nil
}
