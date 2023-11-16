package isanagram

import "strings"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for idx := 0; idx < len(s); idx++ {
		freq[s[idx]-'a']++
		freq[t[idx]-'a']--
	}

	for idx := 0; idx < len(freq); idx++ {
		if freq[idx] != 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	// If the lengths of the strings are different, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Create a map to count the frequency of each character in s
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}

	// Decrease the count for each character found in t
	for _, char := range t {
		if !strings.ContainsAny(t, string(char)) {
			return false
		}
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}
