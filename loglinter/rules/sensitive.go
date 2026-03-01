package rules

import (
	"regexp"
	"strings"

	"github.com/FYNJYserty/log-linter-go/loglinter/config"
)

// Проверка наличия чувствительных данных в сообщении
func ContainsSensitiveData(message string) bool {
	lowerMsg := strings.ToLower(message)

	sensitiveKeywords := config.GetSensitiveKeywords()
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lowerMsg, keyword) {
			return true
		}
	}

	sensitivePatterns := config.GetPatterns()

	for _, pattern := range sensitivePatterns {
		regex := regexp.MustCompile(pattern)
		if regex.MatchString(message) {
			return true
		}
	}

	return false
}
