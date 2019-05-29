package card

import "regexp"

// ValidateName checks if a name of a card is correct
func ValidateName(text string) (b bool) {
	if text == "" {
		return false
	}
	// allow letters, numbers and character "-"
	if ok, _ := regexp.MatchString("^[a-zA-Z\\s0-9-]{4,30}$", text); !ok {
		return false
	}
	return true
}
