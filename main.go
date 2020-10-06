package validator

import (
	"github.com/gobardofw/translator"
)

// NewValidator create new validator
func NewValidator(t translator.Translator) Validator {
	v := new(validatorDriver)
	v.init(t)
	return v
}
