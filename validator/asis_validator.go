package validator

//Asisvalidator compares two strings as they are original
type Asisvalidator struct {
}

//NewAsisvalidator creates a new validator that compares two strings
func NewAsisvalidator() Validator {
	return &Asisvalidator{}
}

//Compare returns the boolean value of the string comparison
func (v *Asisvalidator) Compare(word1 string, word2 string) bool {
	return word1 == word2
}
