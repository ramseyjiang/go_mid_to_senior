package romannum

import "testing"

func TestTerminalExpression(t *testing.T) {
	terminalExpr := TerminalExpression{"I"}

	testCases := []struct {
		context      string
		expectedBool bool
	}{
		{"I", true},
		{"V", false},
		{"XII", true},
		{"IX", true},
		{"", false},
	}

	for _, testCase := range testCases {
		result := terminalExpr.Interpret(testCase.context)

		if result != testCase.expectedBool {
			t.Errorf("TerminalExpression.Interpret() with context '%s' = %v, expected %v", testCase.context, result, testCase.expectedBool)
		}
	}
}

func TestOrExpression(t *testing.T) {
	isOne := &TerminalExpression{"I"}
	isFive := &TerminalExpression{"V"}
	isTen := &TerminalExpression{"X"}

	isOneOrFiveOrTen := &OrExpression{
		Expr1: isOne,
		Expr2: &OrExpression{
			Expr1: isFive,
			Expr2: isTen,
		},
	}

	testCases := []struct {
		context      string
		expectedBool bool
	}{
		{"I", true},
		{"V", true},
		{"X", true},
		{"A", false},
		{"", false},
		{"XXIV", true},
	}

	for _, testCase := range testCases {
		result := isOneOrFiveOrTen.Interpret(testCase.context)

		if result != testCase.expectedBool {
			t.Errorf("OrExpression.Interpret() with context '%s' = %v, expected %v", testCase.context, result, testCase.expectedBool)
		}
	}
}
