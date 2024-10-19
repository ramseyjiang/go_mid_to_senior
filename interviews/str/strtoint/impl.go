package strtoint

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	// Step 1: Trim leading whitespace
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// Step 2: Check for sign
	i := 0
	sign := 1
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3: Process digits
	result := 0
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		// Step 4
		digit := int(s[i] - '0')
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + digit
		i++
	}

	// Apply sign
	return sign * result
}

func myAtoi2(s string) int {
	// Step 1
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// Step 2
	i := 0
	sign := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Step 3
	result := 0
	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		// Step 4
		result = result*10 + int(s[i]-'0')
		if result*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if result*sign < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}

	return result * sign
}
