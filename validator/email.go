package validator

import "regexp"

var emailRegexp = regexp.MustCompile("^[a-zA-z0-9.!#$%&'*+/?^_`{|}~]+@[a-zA-z0-9](?:[a-zA-z0-9]{0," +
	"61}[a-zA-z0-9])?(?:\\.[a-zA-z0-9](?:[a-zA-z0-9]{0,61}[a-zA-z0-9])?)*$")

func (v *Validator) IsEmail(field, email string) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if !emailRegexp.MatchString(email) {
		v.Errors[field] = "not a valid email"
		return false
	}

	return true
}
