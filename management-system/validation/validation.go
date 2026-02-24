package validation

import "regexp"

func IsValidPhone(phoneNo string) bool {
	//valid nos -> 123-456-7890, 123 456 7890, 1234567890
	phoneNoPattern := `\d{3}[- ]?\d{3}[- ]?\d{4}` // [- ]? means either - or space or nothing
	matched, err := regexp.MatchString(phoneNoPattern, phoneNo)
	if err != nil {
		return false
	}
	return matched

}

func IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return false
	}
	return matched
}
