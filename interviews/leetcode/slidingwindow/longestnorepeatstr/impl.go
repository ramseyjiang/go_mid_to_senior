package longestnorepeatstr

func findLongestNoRepeatStr(s string) int {
	charIndex := make(map[byte]int)
	maxLength, start := 0, 0

	for i := 0; i < len(s); i++ {
		if idx, found := charIndex[s[i]]; found && idx >= start {
			start = idx + 1
		}
		charIndex[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}
