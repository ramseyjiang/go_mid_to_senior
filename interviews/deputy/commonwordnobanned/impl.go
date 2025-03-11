package commonwordnobanned

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	if len(paragraph) == 0 {
		return ""
	}

	bannedSet := make(map[string]bool)
	for _, v := range banned {
		bannedSet[strings.ToLower(v)] = true
	}

	// Use regular expression to find all words in the paragraph and ignore punctuation
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(paragraph, -1)

	freq := make(map[string]int)
	for _, word := range words {
		low := strings.ToLower(word)
		freq[low]++
	}

	result, maxCount := "", 0
	for word, count := range freq {
		if !bannedSet[word] && count > maxCount {
			maxCount = count
			result = word
		}
	}

	return result
}
