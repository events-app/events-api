package user

import "testing"

func TestValidateUsername(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"miroslawwalczak@gmail.com", false},
		{"mirek", true},
		{"adm", false},
		{"admin", true},
		{"mmmmmmmmmmmmmmmmmmmm", false},
		{"aa.aa.aa", false},
		{"#username", false},
	}
	for _, test := range tests {
		if got := ValidateUsername(test.input); got != test.want {
			if !got {
				t.Errorf("Username (%q) is invalid but found valid", test.input)
			} else {
				t.Errorf("Username (%q) is valid but found invalid", test.input)
			}
		}
	}
}
