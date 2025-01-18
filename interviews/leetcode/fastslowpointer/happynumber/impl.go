package happynumber

// isHappy determines if a number is a happy number using the fast and slow pointer pattern.
func isHappy(n int) bool {
	getNext := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	slow, fast := n, getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

// isHappy checks whether a given number is a happy number.
// It uses Floyd’s Cycle Detection Algorithm
func isHappy2(n int) bool {
	seen := make(map[int]bool)

	// Helper function to calculate the sum of squares of digits.
	sumSquares := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumSquares(n)
	}

	return n == 1
}
