package rules

import "testing"

func TestContainsSensitiveData(t *testing.T) {
	testCases := []struct {
		msg      string
		expected bool
	}{
		// Отрицательное тестирование - содержат чувствительные данные
		{"user password admin123", true},
		{"api_key=sk_live_abcd1234", true},
		{"access_token: eyJhbGciOiJIUzI1NiIs", true},
		{"secret_key=my_secret_value", true},
		{"credentials: username:pass", true},
		{"private_key=rsa_key_here", true},
		{"refresh_token=abc123def456", true},
		{"bearer authentication failed", true},
		{"aws_secret=wJalrXUtnFEMI", true},
		{"db_password configured", true},
		{"apikey is required", true},

		// Позитивное тестирование - не содержат чувствительных данных
		{"user authenticated successfully", false},
		{"connection established", false},
		{"server started", false},
		{"request completed", false},
		{"database query executed", false},
		{"token validated", false},
		{"token: 1233qwee", true},
		{"authorization header provided", false},
		{"api request successful", false},
	}

	for _, tc := range testCases {
		if result := ContainsSensitiveData(tc.msg); result != tc.expected {
			t.Errorf("ContainsSensitiveData(%q) = %v; expected %v", tc.msg, result, tc.expected)
		}
	}
}
