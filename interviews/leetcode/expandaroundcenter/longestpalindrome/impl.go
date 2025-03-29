package longestPalindrome

// longestPalindrome finds the longest palindromic substring in s.
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLength := 0, 1

	// Helper function to expand around the center and check for palindromes.
	// This approach covers all possible centers, ensuring that the longest palindromic substring is found.
	expandAroundCenter := func(left int, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		// Check for the length of i palindromes.
		// Expands from s[i] as the center, searching for palindromes like “aba” where the length is odd.
		expandAroundCenter(i, i)
		// Check for the length of i+1 palindromes.
		// Expands between s[i] and s[i+1], checking for even-length palindromes like “abba.”
		expandAroundCenter(i, i+1)
	}

	return s[start : start+maxLength]
}
