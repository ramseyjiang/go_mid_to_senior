package crawlerlogfolder

func MinOperations(logs []string) int {
	if len(logs) == 0 {
		return 0
	}

	result := 0
	for _, log := range logs {
		if log == "../" && result <= 0 || log == "./" {
			continue
		}

		if log == "../" {
			result--
		} else {
			result++
		}
	}

	return result
}

func MinOperations2(logs []string) int {
	stack := make([]string, 0)

	for i := range logs {
		switch logs[i] {
		case "../":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case "./":
		default:
			stack = append(stack, logs[i])
		}
	}

	return len(stack)
}
