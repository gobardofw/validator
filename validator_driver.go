package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gobardofw/translator"
)

type validatorDriver struct {
	V *validator.Validate
	T translator.Translator
}

func (v *validatorDriver) init(t translator.Translator) {
	v.V = validator.New()
	v.T = t
}

func (v *validatorDriver) proccessStructValidation(s interface{}, err interface{}, locale string) map[string]string {
	res := make(map[string]string)
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errors {
			// Generate placeholders
			title := suffixedTagOrFallback(s, e.StructField(), "vTitle", locale, e.Field())
			param := e.Param()
			if _, ok := parseFieldTag(s, e.StructField(), "vFormat"); ok {
				param = formatNumericParam(param)
			}
			res[e.Field()] = v.TranslateStruct(s, locale, e.Tag(), map[string]string{
				"field": title,
				"param": param,
			})
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}

// Validator get original validator instance
func (v *validatorDriver) Validator() *validator.Validate {
	return v.V
}

// AddValidation add new validator
func (v *validatorDriver) AddValidation(tag string, val validator.Func) {
	v.V.RegisterValidation(tag, val)
}

// AddTranslation register new translation message to validator translator
func (v *validatorDriver) AddTranslation(locale string, key string, message string) {
	v.T.Register(locale, key, message)
}

// Translate generate translation
func (v *validatorDriver) Translate(locale string, key string, placeholders map[string]string) string {
	return v.T.Translate(locale, key, placeholders)
}

// TranslateStruct generate translation for struct
func (v *validatorDriver) TranslateStruct(s interface{}, locale string, key string, placeholders map[string]string) string {
	return v.T.TranslateStruct(s, locale, key, placeholders)
}

// Struct validates a structs exposed fields, and automatically validates nested structs, unless otherwise specified.
// Return translated errors list on fails
// Return nil on no error
// use vTitle tag for define tag name, (use vTitle_locale) for translation
func (v *validatorDriver) Struct(locale string, s interface{}) map[string]string {
	err := v.V.Struct(s)
	if err != nil {
		return v.proccessStructValidation(s, err, locale)
	}
	return nil
}

// StructExcept validates all fields except the ones passed in.
// Return translated errors list on fails
// Return nil on no error
func (v *validatorDriver) StructExcept(locale string, s interface{}, fields ...string) map[string]string {
	err := v.V.StructExcept(s, fields...)
	if err != nil {
		return v.proccessStructValidation(s, err, locale)
	}
	return nil
}

// StructPartial validates the fields passed in only, ignoring all others.
// Return translated errors list on fails
// Return nil on no error
func (v *validatorDriver) StructPartial(locale string, s interface{}, fields ...string) map[string]string {
	err := v.V.StructPartial(s, fields...)
	if err != nil {
		return v.proccessStructValidation(s, err, locale)
	}
	return nil
}

// Var validates a single variable using tag style validation.
// Return translated errors list on fails
// Return nil on no error
func (v *validatorDriver) Var(locale string, field interface{}, tag string, overrides map[string]string) map[string]string {
	err := v.V.Var(field, tag)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			e := errors[0]
			if msg, ok := overrides[e.Tag()]; ok {
				return map[string]string{e.Field(): msg}
			} else {
				return map[string]string{e.Field(): v.T.Translate(locale, e.Tag(), map[string]string{
					"field": e.Field(),
					"param": e.Param(),
				})}
			}
		}
	}
	return nil
}

// VarWithValue validates a single variable, against another variable/field's value using tag style validation
// Return translated errors list on fails
// Return nil on no error
func (v *validatorDriver) VarWithValue(locale string, field interface{}, other interface{}, tag string, overrides map[string]string) map[string]string {
	err := v.V.VarWithValue(field, other, tag)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			e := errors[0]
			if msg, ok := overrides[e.Tag()]; ok {
				return map[string]string{e.Field(): msg}
			} else {
				return map[string]string{e.Field(): v.T.Translate(locale, e.Tag(), map[string]string{
					"field": e.Field(),
					"param": e.Param(),
				})}
			}
		}
	}
	return nil
}

// Failure make translated failed error for field
func (v *validatorDriver) Failure(field string, message string) map[string]string {
	return map[string]string{field: message}
}
