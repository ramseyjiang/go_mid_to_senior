package reversestr

func ReverseStr(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseStr2(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		j--
		i++
	}
}
