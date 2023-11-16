package firstunique

func firstUniqueChar(s string) int {
	charCount := [26]int{}

	for i := range s {
		charCount[s[i]-'a']++
	}

	for k := range s {
		if charCount[s[k]-'a'] == 1 {
			return k
		}
	}

	return -1
}

func firstUniqueChar2(s string) int {
	charCount := make(map[rune]int)

	// First pass: count the frequency of each character
	for _, char := range s {
		charCount[char]++
	}

	// Second pass: find the first unique character
	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}

	return -1
}
