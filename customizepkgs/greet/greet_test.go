package greet

import (
	"testing"
)

/**
unarygrpc % go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/ramseyjiang/go_mid_to_senior/customizepkgs/unarygrpc 0.009s
*/

/**
unarygrpc % go test
PASS
ok      github.com/ramseyjiang/go_mid_to_senior/customizepkgs/unarygrpc 0.006s

*/

func TestSayHi(t *testing.T) {
	result := SayHi()
	want := "Hey, good morning, welcome to local package."

	if result != want {
		t.Errorf(`Failure, SayHi() = %q, want is: %v`, want, result)
	} else {
		t.Logf("Success, SayHi(), want \"%s\", got \"%s\", it passed", want, result)
	}
}
