package validator

import (
	"strings"
)

//LowerCaseValidator compares two string in lowercase
type LowerCaseValidator struct {
}

//NewLowerCaseValidator creates a new validator that compares two string in low case
func NewLowerCaseValidator() Validator {
	return &LowerCaseValidator{}
}

//Compare returns the boolean values of the lowercase string comparisson
func (v *LowerCaseValidator) Compare(word1 string, word2 string) bool {
	return strings.ToLower(word1) == strings.ToLower(word2)
}
