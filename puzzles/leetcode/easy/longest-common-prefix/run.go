package longest_common_prefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(shortest) > len(strs[i]) {
			shortest = strs[i]
			strs[i] = strs[0]
			strs[0] = shortest
		}
	}

	res := ""
	stop := false
	for i := 0; i < len(shortest); i++ {
		res += string(shortest[i])
		for j := 1; j < len(strs); j++ {
			if strings.Index(strs[j], res) != 0 {
				stop = true
				break
			}
		}

		if stop == true {
			res = res[0 : len(res)-1]
			break
		}
	}

	return res
}
