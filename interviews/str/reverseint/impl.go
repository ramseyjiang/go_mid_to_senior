package reverseint

import "math"

func reverseInt2(x int) int {
	var reversed int
	for x != 0 {
		d := x % 10
		reversed = reversed*10 + d
		x /= 10
		if reversed > math.MaxInt32 || reversed < math.MinInt32 {
			return 0
		}

	}
	return reversed
}

func reverseInt(x int) int {
	var reversed int
	for x != 0 {
		pop := x % 10
		x /= 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && pop < -8) {
			return 0
		}
		reversed = reversed*10 + pop
	}
	return reversed
}
