package reorganisestr

import "sort"

func reorganizeString(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	maxCount := 0
	for _, count := range freq {
		if count > maxCount {
			maxCount = count
		}
	}

	// Characters are sorted by their frequency of any character must not exceed (n+1)/2 to ensure
	// rearrangement is possible.
	if maxCount > (n+1)/2 {
		return ""
	}

	type charCount struct {
		char  rune
		count int
	}

	var chars []charCount
	for char, count := range freq {
		chars = append(chars, charCount{char, count})
	}
	// sorting the characters by frequency in descending order
	sort.Slice(chars, func(i, j int) bool {
		return chars[i].count > chars[j].count
	})

	res := make([]rune, n)
	idx := 0
	// Characters are placed starting from even indices. When all even indices are filled,
	// the placement continues from odd indices.
	for _, cc := range chars {
		for i := 0; i < cc.count; i++ {
			res[idx] = cc.char
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return string(res)
}
