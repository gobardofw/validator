package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gobardofw/validator"
)

func jalaliValidation(fl validator.FieldLevel) bool {
	return v.IsJDate(fl.Field().String())
}

// RegisterJalaliValidation register validations with translations
func RegisterJalaliValidation(val v.Validator) {
	val.AddValidation("jalali", identifierValidation)
	val.AddTranslation("en", "jalali", "Must be a valid jalali date")
	val.AddTranslation("fa", "jalali", "تاریخ وارد شده معتبر نیست")
}
