package convertnuminwords

func ConvertNumInWords(num int) string {
	units := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	if num < 0 || num > 100 {
		return ""
	}

	if num == 100 {
		return "One Hundred"
	}

	if num < 20 {
		return units[num]
	} else {
		if num%10 == 0 {
			return tens[num/10]
		}
		return tens[num/10] + " " + units[num%10]
	}
}
