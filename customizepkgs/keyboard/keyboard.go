package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputFromCmd() (str string) {
	/*
		read (receive and store) input from the program’s standard input, which all keyboard input goes to
		NewReader will return a new bufio.Reader
		os.Stdin store the standard input from the keyboard.
	*/
	reader := bufio.NewReader(os.Stdin)

	/*
		Return everything the user has typed, up to where they pressed the Enter key.
		ReadString will return the user typed, as a string.
		The ReadString method requires an argument with a rune (character) that marks the end of the input.
		'\n' means everything up until the newline rune will be read.
		It should not be named error, because it is the go keyword. That's why named err.
	*/
	// all inputs from cmd are all string
	input, _ := reader.ReadString('\n')

	// TrimSpace is used to remove all whitespace characters (newlines, tabs, and regular spaces) from the start and end of a string. from the input string
	return strings.TrimSpace(input)
}

/*
	GetFloat is used to get inputs from the keyboard.
	Receive the input from command line and convert it to float.
	Return float64 and error.
*/
func GetFloat() (number float64, err error) {
	// ParseFloat is used to convert the string to a number, and returns it as a float64 value.
	number, err = strconv.ParseFloat(getInputFromCmd(), 64)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return number, nil
}
