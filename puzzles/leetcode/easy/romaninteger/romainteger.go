package romaninteger

/*
Runtime: 12 ms, faster than 62.86% of Go online submissions for roman to integer.
Memory Usage: 3.3 MB, less than 16.9% of Go online submissions for roman to integer.
*/
func romanToInt1(s string) int {
	romanArr := [7]int{1, 5, 10, 50, 100, 500, 1000}
	arr := make(map[int]int)
	sum := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "I":
			arr[i] = romanArr[0]
		case "V":
			arr[i] = romanArr[1]
		case "X":
			arr[i] = romanArr[2]
		case "L":
			arr[i] = romanArr[3]
		case "C":
			arr[i] = romanArr[4]
		case "D":
			arr[i] = romanArr[5]
		case "M":
			arr[i] = romanArr[6]
		}

		if arr[i-1] < arr[i] { // Because arr[i-1] has been added in the sum, that's why here is sub arr[i-1]*2.
			sum += arr[i] - arr[i-1]*2
		} else {
			sum += arr[i]
		}
	}

	return sum
}

/*
The best is
Runtime: 8 ms, faster than 83.54% of Go online submissions for roman to integer.
Memory Usage: 2.9 MB, less than 15.76% of Go online submissions for roman to integer.
*/
func romanToInt2(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romans[string(s[i])] < romans[string(s[i+1])] {
			result -= romans[string(s[i])]
		} else {
			result += romans[string(s[i])]
		}
	}
	return result
}
