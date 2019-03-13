package user

import "regexp"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// get users from DB
var Users = []User{
	User{ID: 1, Username: "admin", Password: "admin", Role: "ADMIN"},
	User{ID: 2, Username: "user", Password: "user", Role: "USER"},
}

// ValidateUsername checks if username has correct structure
// lowercase letters, uppercase letters, numbers, minimal length 4, maximum length 16
func ValidateUsername(text string) (b bool) {
	if text == "" {
		return false
	}
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", text); !ok {
		return false
	}
	return true
}
