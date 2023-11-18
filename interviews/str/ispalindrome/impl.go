package ispalindrome

import (
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var builder strings.Builder
	builder.Grow(len(s))

	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}

	str := builder.String()

	// Check if the string is a palindrome
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}

	pattern := "[^a-zA-Z0-9]"
	s = strings.ToLower(regexp.MustCompile(pattern).ReplaceAllString(s, ""))

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
