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
