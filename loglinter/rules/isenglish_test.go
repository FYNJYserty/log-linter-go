package rules

import "testing"

func TestIsEnglish(t *testing.T) {
	testCases := []testCase{
		{"This is an English message.", true},
		{"Сообщение на русском языке", false},
		{"Message with numbers 12345", true},
		{"Message_with_underscore=42", true},
		{"Message-with-dash:ok", true},
		{"Message with special characters !@#$%^&*()", true},
		{"Mixed message: English and русский.", false},
		{"server started! ", true},
		{"server started! 🚀", true},
	}

	for _, tc := range testCases {
		if result := IsEnglish(tc.msg); result != tc.expected {
			t.Errorf("IsEnglish(%q) = %v; expected %v", tc.msg, result, tc.expected)
		}
	}
}
