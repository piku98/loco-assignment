package utils

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

func NumericString(fl validator.FieldLevel) bool {
	_, err := strconv.Atoi(fl.Field().String())
	return err == nil
}
