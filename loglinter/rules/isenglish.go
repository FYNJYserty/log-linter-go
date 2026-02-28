package rules

import "regexp"

func IsEnglish(message string) bool {
	var isEnglishRegex = regexp.MustCompile(`^[\p{Latin}\p{N}\p{Zs}\p{P}\p{S}]+$`)
	return isEnglishRegex.MatchString(message)
}
