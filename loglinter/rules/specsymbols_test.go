package rules

import "testing"

func TestIsNoSpecSymbols(t *testing.T) {
	testCases := []testCase{
		{"This is a normal message.", false},
		{"Message with emoji 🚀", false},
		{"Message with special characters !@#$%^&*()", false},
		{"Another normal text", true},
		{"server started! 🚀", false},
		{"warning: something went wrong...", false},
	}

	for _, tc := range testCases {
		if result := IsNoSpecSymbols(tc.msg); result != tc.expected {
			t.Errorf("IsNoSpecSymbols(%q) = %v; expected %v", tc.msg, result, tc.expected)
		}
	}
}
