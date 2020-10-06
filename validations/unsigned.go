package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gobardofw/validator"
)

func unsignedValidation(fl validator.FieldLevel) bool {
	return v.IsUnsigned(fl.Field().String())
}

// RegisterUnsignedValidation register validations with translations
func RegisterUnsignedValidation(val v.Validator) {
	val.AddValidation("unsigned", identifierValidation)
	val.AddTranslation("en", "unsigned", "Must be a unsigned number")
	val.AddTranslation("fa", "unsigned", "باید یک عدد صحیح مثبت باشد")
}
