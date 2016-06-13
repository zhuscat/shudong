package utils

import (
	"regexp"
)

func IsValidUsername(username string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]{3,16}$", username)
	return match
}

func IsValidPassword(password string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]{6,20}$", password)
	return match
}

func IsValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^\S+@\S+\.\S+$`, email)
	return match
}
