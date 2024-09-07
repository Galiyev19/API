package validator

import "regexp"

type Validator struct {
	Errors map[string]string
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Valid return true if errors map is empty
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// Add errors in map
func (v *Validator) AddError(key, message string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Check email is valid
func (v *Validator) Matches(email string, regX *regexp.Regexp) bool {
	return regX.MatchString(email)
}
