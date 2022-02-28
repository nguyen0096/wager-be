package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// GetValidate validate struct fields tag
func Get() *validator.Validate {
	return validate
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "gt":
		return fmt.Sprintf("this field must be larger than %s", fe.Param())
	case "gte":
		return fmt.Sprintf("this field must be larger or equal %s", fe.Param())
	case "lte":
		return fmt.Sprintf("this field must be smaller or equal %s", fe.Param())
	}
	return fe.Error()
}
