package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gobardofw/validator"
)

func telValidation(fl validator.FieldLevel) bool {
	return v.IsTel(fl.Field().String())
}

// RegisterTelValidation register validations with translations
func RegisterTelValidation(val v.Validator) {
	val.AddValidation("tel", identifierValidation)
	val.AddTranslation("en", "tel", "Must be a valid  tel")
	val.AddTranslation("fa", "tel", "شناسه وارد شده معتبر نیست")
}
