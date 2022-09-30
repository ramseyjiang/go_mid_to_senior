package palindromenum

import (
	"strconv"
)

/*
Runtime: 11-28 ms, faster than 67.88% - 98.5% of Go online submissions for palindrome num.
Memory Usage: 4.8 MB, less than 48.2% of Go online submissions for palindrome num.
*/
func isPalindrome1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)
	bytes := []rune(str)

	for from, to := 0, len(str)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}

	return string(bytes) == str
}

// Golang does not have built-in methods for string reversal.
// Cannot convert directly from a slice of runes to an integer either.
// Hence, convert the runes to a string, then convert the string to an int.
/*
Runtime: 7-28 ms, faster than 67.88% - 99.7% of Go online submissions for palindrome num.
Memory Usage: 4.8 MB, less than 48.2% of Go online submissions for palindrome num.
*/
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// convert int to slice of runes
	runes := []rune(strconv.Itoa(x))

	// reverse the runes in the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// convert back to int, via string
	y, err := strconv.Atoi(string(runes))

	// check for an error (overflow) and equality
	if err == nil && x == y {
		return true
	}
	return false
}
