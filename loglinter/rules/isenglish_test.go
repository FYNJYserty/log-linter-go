package rules

import "testing"

func TestIsEnglish(t *testing.T) {
	testCases := []testCase{
		{"This is an English message.", false},
		{"Сообщение на русском языке", false},
		{"Message with numbers 12345", true},
		{"Message with special characters !@#$%^&*()", false},
		{"Mixed message: English and русский.", false},
		{"server started! 🚀", false},
	}

	for _, tc := range testCases {
		if result := IsEnglish(tc.msg); result != tc.expected {
			t.Errorf("IsEnglish(%q) = %v; expected %v", tc.msg, result, tc.expected)
		}
	}
}
