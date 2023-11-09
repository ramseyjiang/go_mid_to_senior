package rop

func removeOuterParentheses(s string) string {
	var result []rune
	var balance int

	for _, char := range s {
		if char == '(' {
			// If balance is greater than 0, this is not an outer parenthesis.
			if balance > 0 {
				result = append(result, char)
			}
			balance++
		} else if char == ')' {
			balance--
			// If balance is greater than 0, this is not an outer parenthesis.
			if balance > 0 {
				result = append(result, char)
			}
		}
	}

	return string(result)
}
