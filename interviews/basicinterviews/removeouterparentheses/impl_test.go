package rop

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect string
	}{
		{"Empty string", "", ""},
		{"No change needed", "(())", "()"},
		{"Single pair", "()()", ""},
		{"Nested parentheses", "((()))", "(())"},
		{"Multiple sets", "(()())(())", "()()()"},
		{"Complex nesting", "(()(()))", "()(())"},
		{"All outer", "(()())(())(()(()))", "()()()()(())"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := removeOuterParentheses(tc.input)
			if result != tc.expect {
				t.Errorf("removeOuterParentheses(%q) = %q, expected %q", tc.input, result, tc.expect)
			}
		})
	}
}
