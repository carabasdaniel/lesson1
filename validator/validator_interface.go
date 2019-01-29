package validator

//Validator interface does something
type Validator interface {
	Compare(firstWord string, secondWord string) bool
}
