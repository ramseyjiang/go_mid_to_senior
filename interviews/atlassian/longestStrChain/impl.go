package longestStrChain

import (
	"sort"
)

func longestStrChan(words []string) int {
	// sort the string slice from short to long
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	longestChain := make(map[string]int)
	maxLen := 0

	for _, word := range words {
		longestChain[word] = 1
		for i := 0; i < len(word); i++ {
			// Concatenates two substrings, without the character at position i.
			// It is used to check whether the prev string already in the longest chain.
			prev := word[:i] + word[i+1:]

			// if the prev has already in, the longest chain length adds 1.
			if chainLen, exists := longestChain[prev]; exists && chainLen+1 > longestChain[word] {
				longestChain[word] = chainLen + 1
			}
		}

		// compare the longestChain[word] length with the previous maxLen.
		if longestChain[word] > maxLen {
			maxLen = longestChain[word]
		}
	}

	return maxLen
}
