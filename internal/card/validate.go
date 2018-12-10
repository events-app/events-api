package card

import "regexp"

func ValidateName(text string) (b bool) {
	if text == "" {
		return false
	}
	// allow letters, numbers and character "-"
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9-]{4,30}$", text); !ok {
		return false
	}
	return true
}
