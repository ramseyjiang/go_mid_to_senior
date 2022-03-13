package puzzles

import "fmt"

/**
func first(sourceArray []string) {
	var destArray []*string

	for _, element := range sourceArray {
		destArray = append(destArray, &element)
	}

	// Trying to print elements of destArray
	for _, element := range destArray {
		println(*element)
	}
}
*/

func TriggerPointer() {
	sourceArray := []string{"first", "second", "third"}
	/**
	first(sourceArray)
	*/
	second(sourceArray)
	third(sourceArray)
}

func second(sourceArray []string) {
	for i, elem := range sourceArray {
		fmt.Printf("Index: %d, element: %s, pointer: %p \n", i, elem, &elem)
	}
}

func third(sourceArray []string) {
	for i, elem := range sourceArray {
		fmt.Printf("Index: %d, element: %s, pointer: %p \n", i, elem, &sourceArray[i])
	}
}

// In the second, in the loop range, it prints the address only from elem.
