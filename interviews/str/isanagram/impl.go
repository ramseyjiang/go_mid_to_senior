package isanagram

import (
	"reflect"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int)

	for i, _ := range s {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram4(s string, t string) bool {
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

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := [26]int{}
	tArr := [26]int{}

	for i := range s {
		sArr[s[i]-'a']++
	}

	for j := range t {
		tArr[t[j]-'a']++
	}

	if !reflect.DeepEqual(sArr, tArr) {
		return false
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
		count[char]--
		if count[char] < 0 {
			return false
		}
	}

	return true
}

func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	sMap := make(map[string]int)
	for i, _ := range sArr {
		sMap[sArr[i]] += 1
		sMap[tArr[i]] -= 1
	}
	for _, val := range sMap {
		if val != 0 {
			return false
		}
	}
	return true
}
