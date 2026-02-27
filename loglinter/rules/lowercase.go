package rules

import (
	"unicode"
	"unicode/utf8"
)

// Проврека на то, что сообщение начинается с заглавной буквы
func CheckLowerCase(message string) bool {
	if len(message) == 0 {
		return true
	}

	r, _ := utf8.DecodeRuneInString(message)

	if unicode.IsUpper(r) {
		return false
	}
	return true
}
