package rules

import (
	"regexp"

	"github.com/forPelevin/gomoji"
)

// Проверка наличия специальных символов/эмодзи в сообщении
func IsNoSpecSymbols(message string) bool {
	var specialSymbolsRegex = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>=]`)
	if gomoji.ContainsEmoji(message) || specialSymbolsRegex.MatchString(message) {
		return false
	}
	return true
}
