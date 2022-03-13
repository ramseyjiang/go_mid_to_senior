package joinstr

import (
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		// If the slice has just two items, just join them together with "and".
		return phrases[0] + " and " + phrases[1]
	} else { // Otherwise, use the same code we always have.
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}
