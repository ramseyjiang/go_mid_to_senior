package strtoint

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	// Step 1: Trim leading whitespace
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Step 2: Check for sign
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		start++
	}

	result := 0

	// Step 3: Process digits
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		digit := int(s[i] - '0')

		// Step 4 & 5: Convert to integer and handle overflow
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		result = result*10 + digit
	}

	// Apply sign
	return sign * result
}

func myAtoi2(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Check for sign at the start
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		start++
	}

	var builder strings.Builder

	// Process only the numeric part
	for i := start; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			builder.WriteRune(rune(s[i]))
		} else {
			break
		}
	}

	str := builder.String()

	// Handle empty string or string with only '+' or '-'
	if str == "" {
		return 0
	}

	result, err := strconv.Atoi(str)
	if err != nil {
		// Handle integer overflow
		if sign == 1 {
			return 1<<31 - 1
		}
		return -1 << 31
	}

	// Apply sign and clamp within 32-bit integer range
	result *= sign
	if result > 1<<31-1 {
		return 1<<31 - 1
	} else if result < -1<<31 {
		return -1 << 31
	}

	return result
}
