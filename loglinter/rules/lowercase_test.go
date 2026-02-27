package rules

import "testing"

func TestCheckLowerCase(t *testing.T) {
	cases := []testCase{
		{"hello", true},
		{"Hello", false},
		{"", true},
	}

	for _, c := range cases {
		if got := CheckLowerCase(c.msg); got != c.expected {
			t.Errorf("CheckLowerCase(%q) = %v, want %v", c.msg, got, c.expected)
		}
	}
}
