package longestsubstrnorepeat

func findLongestSubstr(s string) int {
	maxLength, start := 0, 0
	charIndex := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		// After a repeating character is found, the if condition will work.
		if idx, found := charIndex[s[i]]; found && idx >= start {
			start = idx + 1
		}

		// If no repeating characters found, element always fills in charIndex.
		// If the repeating found, the repeating index will refresh the old index.
		charIndex[s[i]] = i

		// Compare substrings and get the longest substring
		// Because i starts from 0, it should use i-start+1, not use i-start
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}
