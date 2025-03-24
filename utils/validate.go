package utils

import "github.com/go-playground/validator"

var validate = validator.New()

func ValidateStruct(s interface{}) []string {
	err := validate.Struct(s)
	if err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" validation failed: "+err.ActualTag())
		}
		return errors
	}
	return nil
}
