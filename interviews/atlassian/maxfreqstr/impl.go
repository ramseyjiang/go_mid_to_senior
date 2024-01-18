package maxfreqstr

func maxFreqStr(s string, maxLetters int, minSize int, maxSize int) int {
	subs := map[string]int{}
	maxFreq := 0

	for i := 0; i <= len(s)-minSize; i++ {
		substr := s[i : i+minSize]
		ch := map[rune]int{}
		for _, x := range substr {
			ch[x]++
		}
		if len(ch) <= maxLetters && len(ch) <= maxSize {
			subs[substr]++
			maxFreq = max(maxFreq, subs[substr])
		}
	}

	return maxFreq
}

func maxFreqStr2(s string, maxLetters int, minSize int, maxSize int) int {
	count := make(map[string]int)
	maxFreq := 0

	for i := 0; i <= len(s)-minSize; i++ {
		substr := s[i : i+minSize]
		if isValid(substr, maxLetters) {
			count[substr]++
			if count[substr] > maxFreq {
				maxFreq = count[substr]
			}
		}
	}

	return maxFreq
}

func isValid(substr string, maxLetters int) bool {
	uniqueLetters := make(map[rune]int)
	for _, char := range substr {
		uniqueLetters[char]++
		if len(uniqueLetters) > maxLetters {
			return false
		}
	}
	return true
}
