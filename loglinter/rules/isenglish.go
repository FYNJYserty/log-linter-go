package rules

import "regexp"

// Функция проверки на английский язык через регулярные выражения
func IsEnglish(message string) bool {
	var isEnglishRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return isEnglishRegex.MatchString(message)
}
