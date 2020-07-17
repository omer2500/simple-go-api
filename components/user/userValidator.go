package user

import (
	"strings"
	"unicode"
)

func vlidateUserInsert(user User) []string {
	errors := []string{}
	if len(strings.TrimSpace(user.Password)) == 0 || !isValid(user.Password) {
		errors = append(errors, "password is invalid or empty")
	}
	if len(strings.TrimSpace(user.Username)) == 0 {
		errors = append(errors, "password is invalid or empty")
	}
	return errors
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 8 {
		hasMinLen = true
	}
	if len(s) <= 50 {
		hasMaxLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial && hasMaxLen
}
