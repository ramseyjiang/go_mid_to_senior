package joinstr_test

import (
	"fmt"
	"testing"

	"golang_learn/customizepkgs/joinstr"
)

// In comments, it includes three ways to run tests.

// The first way, run tests: go test second_way_test.go
// You can use Errorf to include additional information in your test’s failure messages, such as the
// arguments you passed to a function, the return value you got, and the value you were expecting.
// If tests are not passed, the result is following:
/**
--- FAIL: TestTwoElements (0.00s)				//Function name of failing test
    second_way_test.go:24: no test written yet		//Filaname and line number
--- FAIL: TestThreeElements (0.00s)				//Function name of failing test
    second_way_test.go:34: JoinWithCommas([]string{"apple", "orange", "pear"}) = "apple, orange and pear", want "apple and orange"
    //Filaname and line number
FAIL
FAIL    command-line-arguments  0.006s			//Status for the example_test package overall
FAIL
*/

// If tests are passed, the result is as below.
/**
ok      command-line-arguments  0.007s
*/

// The second way, run all tests is as following:
/**
 go test second_way_test.go -v
=== RUN   TestOneElement
--- PASS: TestOneElement (0.00s)
=== RUN   TestTwoElements
--- PASS: TestTwoElements (0.00s)
=== RUN   TestThreeElements
--- PASS: TestThreeElements (0.00s)
PASS
ok      command-line-arguments  0.008s
*/

// The third way, run all tests is as below:
/**
go test second_way_test.go -v -run Two

=== RUN   TestOneElement
--- PASS: TestOneElement (0.00s)
PASS
ok      command-line-arguments  0.007s
*/

// If the third way, use One to replace two, it will run as following:
/**
=== RUN   TestOneElement
--- PASS: TestOneElement (0.00s)
PASS
ok      command-line-arguments  0.007s
*/

// The fourth way, run all tests is below:
/**
go test second_way_test.go -v -run Elements

=== RUN   TestTwoElements
--- PASS: TestTwoElements (0.00s)
=== RUN   TestThreeElements
--- PASS: TestThreeElements (0.00s)
PASS
ok      command-line-arguments  0.007s
*/

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := joinstr.JoinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	} else {
		t.Logf(successString(list, got, want))
	}
}

// Function name should begin with "Test".
// Name after "Test" can be whatever you want.
// Test functions must accept a single parameter: a pointer to a testing.T value.
func TestTwoElements(t *testing.T) { // Function will be passed a  pointer to a testing.T value.
	list := []string{"apple", "orange"}
	want := "apple and orange"          // want is the return value we want.
	got := joinstr.JoinWithCommas(list) // git is the return value we actually got.
	if got != want {
		// Instead of calling t.Errorf(), call errorStirng(), it's an error helper function.
		t.Error(errorString(list, got, want))
	} else {
		t.Logf(successString(list, got, want))
	}
}

// Test functions must accept a single parameter: a pointer to a testing.T value.
// You can report that a test has failed by calling methods (such as Error) on the testing.T value.
// Most methods accept a string with a message explaining the reason the test failed.
func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"   // want is the return value we want.
	got := joinstr.JoinWithCommas(list) // git is the return value we actually got.

	if got != want {
		// Instead of calling t.Errorf(), call errorStirng(), it's an error helper function.
		t.Error(errorString(list, got, want)) // remove this line, the test will be passed.
	} else {
		t.Logf(successString(list, got, want))
	}
}

// Functions within a _test.go file whose names do not begin with Test are not run by go test.
// They can be used by tests as “helper” functions.
// errorString is an error helper function.
func errorString(list []string, got string, want string) string {
	// Errorf() works similarly to Error, but it accepts a formatting string just like the fmt.Printf function.
	// t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
	return fmt.Sprintf("JoinStr(%#v) = \"%q\", want \"%q\"", list, got, want)
}

func successString(list []string, got string, want string) string {
	return fmt.Sprintf("Success, JoinStr(%#v), want \"%q\", got \"%q\", it passed", list, want, got)
}
