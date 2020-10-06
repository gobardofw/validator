package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gobardofw/validator"
)

func mobileValidation(fl validator.FieldLevel) bool {
	return v.IsMobile(fl.Field().String())
}

// RegisterMobileValidation register validations with translations
func RegisterMobileValidation(val v.Validator) {
	val.AddValidation("mobile", identifierValidation)
	val.AddTranslation("en", "mobile", "Must be a valid mobile")
	val.AddTranslation("fa", "mobile", "شناسه وارد شده معتبر نیست")
}
