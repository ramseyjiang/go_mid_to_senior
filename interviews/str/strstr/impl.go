package strstr

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr2(haystack string, needle string) int {
	foundAt := -1
	nl := len(needle)

	for i := 0; i < (len(haystack) - nl + 1); i++ {
		subs := haystack[i:(i + nl)]
		if subs == needle {
			foundAt = i
			break
		}
	}
	return foundAt
}
