package pin

import (
	"math/rand"
	"strconv"
	"time"
)

const totalNum = 1000

func GeneratePins() []string {
	var res []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < totalNum; i++ {
		var pin string
		// While generating the random number it checks if the generated pin has any two consecutive digits
		// that are the same or three consecutive digits that are incremental, if it does,
		// it doesn't add it to the pin string and generate a new number again.
		for j := 0; j < 4; j++ {
			n := rand.Intn(10)
			if j > 0 && (n == int(pin[j-1]-'0') || (n-1 == int(pin[j-1]-'0')) || (n+1 == int(pin[j-1]-'0'))) {
				j--
				continue
			}
			pin += strconv.Itoa(n)
		}
		res = append(res, pin)
	}
	return res
}

func hasConsecutiveSame(pin string) bool {
	for i := 0; i < 3; i++ {
		if pin[i] == pin[i+1] {
			return true
		}
	}
	return false
}

func hasIncrementalConsecutive(pin string) bool {
	for i := 0; i < 2; i++ {
		if int(pin[i])+1 == int(pin[i+1]) && int(pin[i+1])+1 == int(pin[i+2]) {
			return true
		}
	}
	return false
}
