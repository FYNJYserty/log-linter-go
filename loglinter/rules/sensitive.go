package rules

import (
	"regexp"
	"strings"
)

// Список возможных чувствительных ключевых слов для обнаружения в логах
var sensitiveKeywords = []string{
	"password",
	"passwd",
	// plain "token" removed, rely on pattern matching with punctuation
	"pwd",
	"api_key",
	"apikey",
	"api-key",
	"secret",
	"credential",
	"credentials",
	"private_key",
	"privatekey",
	"private-key",
	"access_token",
	"accesstoken",
	"refresh_token",
	"refreshtoken",
	"bearer",
	"aws_secret",
	"db_password",
	"database_password",
}

// Проверка наличия чувствительных данных в сообщении
func ContainsSensitiveData(message string) bool {
	lowerMsg := strings.ToLower(message)

	for _, keyword := range sensitiveKeywords {
		if strings.Contains(lowerMsg, keyword) {
			return true
		}
	}

	sensitivePatterns := []string{
		// patterns capture keywords followed by ':' or '='. include plain
		// "token" here so that "token: ..." or "token = ..." triggers the
		// detection while "token validated" does not.
		`(?i)(token|password|passwd|pwd|api[_-]?key|secret|credential|access[_-]?token|refresh[_-]?token|bearer|private[_-]?key)\s*[:=]`,
	}

	for _, pattern := range sensitivePatterns {
		regex := regexp.MustCompile(pattern)
		if regex.MatchString(message) {
			return true
		}
	}

	return false
}
