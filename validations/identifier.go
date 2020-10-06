package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gobardofw/validator"
)

func identifierValidation(fl validator.FieldLevel) bool {
	return v.IsIdentifier(fl.Field().String())
}

// RegisterIdentifierValidation register identifier validator and it translations
func RegisterIdentifierValidation(val v.Validator) {
	val.AddValidation("identifier", identifierValidation)
	val.AddTranslation("en", "identifier", "Must be a valid (numeric) identifier")
	val.AddTranslation("fa", "identifier", "شناسه وارد شده معتبر نیست")
}
