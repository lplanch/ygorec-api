package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

func ValidateBanlist(fl validator.FieldLevel) bool {
	println("BONJOUR")
	match, err := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2}", fl.Field().String())
	if err != nil {
		return false
	}
	return match
}

func GoValidator(s interface{}, config []gpc.ErrorMetaConfig) (interface{}, int) {
	var validate *validator.Validate
	validators := gpc.NewValidator(validate)
	bind := gpc.NewBindValidator(validators)

	errResponse, errCount := bind.BindValidator(s, config)
	return errResponse, errCount
}
