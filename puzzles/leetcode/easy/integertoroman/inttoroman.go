package integertoroman

// IntToRoman is using the hard code way.
// thousands  ==> M
// hundreds   ==> CM, D, CD, C
// tens       ==> XC, L, XL, X
// ones       ==> IX, V, IV, I
func IntToRoman(num int) (str string) {
	thousands := [4]string{"", "M", "MM", "MMM"}
	hundreds := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
