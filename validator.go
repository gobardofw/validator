package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator interface
type Validator interface {
	// Validator get original validator instance
	Validator() *validator.Validate
	// AddValidation add new validator
	AddValidation(tag string, v validator.Func)

	// AddTranslation register new translation message to validator translator
	AddTranslation(locale string, key string, message string)
	// Translate generate translation
	Translate(locale string, key string, placeholders map[string]string) string
	// TranslateStruct generate translation for struct
	TranslateStruct(s interface{}, locale string, key string, placeholders map[string]string) string

	// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
	// Return translated errors list on fails
	// Return nil on no error
	// use vTitle tag for define tag name, (use vTitle_locale) for translation
	// use vFormat tag for format parameter as number
	Struct(locale string, s interface{}) map[string]string
	// StructExcept validates all fields except the ones passed in.
	// Return translated errors list on fails
	// Return nil on no error
	StructExcept(locale string, s interface{}, fields ...string) map[string]string
	// StructPartial validates the fields passed in only, ignoring all others.
	// Return translated errors list on fails
	// Return nil on no error
	StructPartial(locale string, s interface{}, fields ...string) map[string]string

	// Var validates a single variable using tag style validation.
	// Return translated errors list on fails
	// Return nil on no error
	Var(locale string, field interface{}, tag string, overrides map[string]string) map[string]string
	// VarWithValue validates a single variable, against another variable/field's value using tag style validation
	// Return translated errors list on fails
	// Return nil on no error
	VarWithValue(locale string, field interface{}, other interface{}, tag string, overrides map[string]string) map[string]string

	// Failure make translated failed error for field
	Failure(field string, message string) map[string]string
}
