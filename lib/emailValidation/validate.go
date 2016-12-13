package emailvalidation

import "regexp"

//Validates that the email provided has the correct format
func Validate(e string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(e)
}
