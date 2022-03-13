package puzzles

import (
	"fmt"
	"os"
)

func TriggerScope() {
	example()
	solution()
}

/** Expected outputs.
there
are
no
strings
on
me
*/
func example() {
	var data []string
	var err error

	killSwitch := os.Getenv("KILLSWITCH")
	if killSwitch == "" {
		fmt.Println("kill switch is off")
		data, err = getData()
		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
	fmt.Println(err)
}

func getData() ([]string, error) {
	return []string{"there", "are", "no", "strings", "on", "me"}, nil
}

func solution() {
	var data []string
	var err error // Declaring err to make sure we can use = instead of :=

	killSwitch := os.Getenv("KILLSWITCH")

	if killSwitch == "" {
		fmt.Println("kill switch is off")
		data, err = getData() // Here replace ":=" by "=", then the data scope is changed.

		if err != nil {
			panic("ERROR!")
		}

		fmt.Printf("Data was fetched! %d\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
	fmt.Println(err)
}

// In the example(), the data and err are defined both at the beginning and in the if condition.
// Outside the if condition is the global data, inside the if condition is the partial data.
// The partial one and the global one won't influence each other.

// In the solution, the data and err are defined only at the beginning, so the data in the if condition also is the global data.
