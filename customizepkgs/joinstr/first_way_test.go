package joinstr_test

import (
	"testing"

	"github.com/ramseyjiang/go_mid_to_senior/customizepkgs/joinstr"
)

/**
joinstr % go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/ramseyjiang/go_mid_to_senior/customizepkgs/joinstr       0.007s
*/

/**
joinstr % go test
PASS
ok      github.com/ramseyjiang/go_mid_to_senior/customizepkgs/joinstr       0.006s
*/

// If the tests in the same package with codes, please run tests "click run button."
// If not, you can run tests using any of the following ways.

// This test is better than second_way_test.go, because it moves common check parts in three different tests outside
// in a same function. Then it will use less time to run all tests anf faster than second_way_test.go.
// The compare result is as following:
/**
% go test second_way_test.go
ok      command-line-arguments  0.007s
% go test first_way_test.go
ok      command-line-arguments  0.006s
*/

// The way in this file it is called Table-driven tests.
// Table-driven tests are tests that process “tables” of inputs and expected outputs. They pass each set of input to
// the code being tested, and check that the code’s output matches the expected values.

/**	The upgrade
 go test first_way_test.go -v
=== RUN   TestJoinStr
    first_way_test.go:41: Success, JoinStr([]string(nil)), want "", got "", it passed
    first_way_test.go:41: Success, JoinStr([]string{"apple"}), want "apple", got "apple", it passed
    first_way_test.go:41: Success, JoinStr([]string{"apple", "orange"}), want "apple and orange", got "apple and orange", it passed
    first_way_test.go:41: Success, JoinStr([]string{"apple", "orange", "pear"}), want "apple, orange, and pear", got "apple, orange, and pear", it passed
--- PASS: TestJoinStr (0.00s)
PASS
ok      command-line-arguments  0.011s

*/

type testData struct {
	list []string // The slice we'll pass to target method.
	want string   // The string we expect return.
}

func TestJoinStr(t *testing.T) {
	tests := []testData{
		{},
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := joinstr.JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		} else {
			t.Logf("Success, JoinStr(%#v), want \"%s\", got \"%s\", it passed", test.list, test.want, got)
		}
	}
}
