package implstrstr

import (
	"strings"
)

func strStr2(hatstack string, needle string) int {
	if strings.Contains(hatstack, needle) {
		return strings.Index(hatstack, needle)
	} else {
		return -1
	}
}

func strStr(haystack string, needle string) int {
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
