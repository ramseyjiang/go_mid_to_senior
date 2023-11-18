package commonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	for i, _ := range strs {
		min := i
		for j := i + 1; j < len(strs); j++ {
			if strs[j] < strs[min] {
				min = j
			}
		}
		strs[i], strs[min] = strs[min], strs[i]
	}
	index := 0
	first := strs[0]
	last := strs[len(strs)-1]
	for i := 0; i < len(first); i++ {
		if first[i] == last[i] {
			index++
		} else {
			break
		}
	}
	return first[:index]
}

func longestCommonPrefix2(strs []string) string {
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
