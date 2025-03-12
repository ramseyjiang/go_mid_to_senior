package calculate2

func calculate(s string) int {
	stack := []int{}
	currentNum := 0
	lastOp := '+'

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			currentNum = currentNum*10 + int(c-'0')
		} else if c == ' ' {
			continue
		} else {
			process(lastOp, &stack, currentNum)
			lastOp = rune(c)
			currentNum = 0
		}
	}

	process(lastOp, &stack, currentNum)
	res := 0
	for _, num := range stack {
		res += num
	}

	return res
}

func process(op rune, stack *[]int, num int) {
	switch op {
	case '+':
		*stack = append(*stack, num)
	case '-':
		*stack = append(*stack, -num)
	case '*':
		prev := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		*stack = append(*stack, prev*num)
	case '/':
		prev := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		res := prev / num
		*stack = append(*stack, res)
	}
}
