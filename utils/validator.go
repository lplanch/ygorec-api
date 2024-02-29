package util

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

/**
*** CUSTOMS VALIDATOR FUNCTIONS
***/

func ValidateBanlist(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) == 0 {
		return true
	}
	match, err := regexp.MatchString("^[0-9]{4}-[0-9]{2}-[0-9]{2}$", fl.Field().String())
	if err != nil {
		return false
	}
	return match
}

/**
*** API ERROR STRINGIFY
**/

type ApiError struct {
	Param   string
	Message string
}

func tagToString(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "number":
		return "Field must be a number"
	case "gt":
		return "Field must be greater than " + fe.Param()
	case "gte":
		return "Field must be greater or equal than " + fe.Param()
	case "banlistdate":
		return "Field must match the banlist date format \"%Y-%m-%d\""
	}
	return fe.Error()
}

func ErrorHandler(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), tagToString(fe)}
		}
		return out
	}
	return []ApiError{}
}

/**
*** INIT & VALIDATE BINDING
**/

func InitValidator() {
	validate = validator.New()
	validate.RegisterValidation("banlistdate", ValidateBanlist)
}

func Validate(s interface{}) []ApiError {
	err := validate.Struct(s)
	if err != nil {
		return ErrorHandler(err)
	}
	return nil
}
