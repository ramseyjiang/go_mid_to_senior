package valid

func isValid(s string) bool {
	// rune is an alias for int32 and is equivalent to int32 in all ways.
	// It is used, by convention, to distinguish character values from integer values.
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if _, exist := mapping[char]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
