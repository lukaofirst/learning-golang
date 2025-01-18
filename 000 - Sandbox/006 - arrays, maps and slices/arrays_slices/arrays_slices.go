package main

import "fmt"

func main() {
	// Creating an array
	names := []string{"lorem", "ipsum", "dolor"}

	// Taking a slice of an array
	sliceStartingAtIndexOne := names[1:] // gets all the elements starting at index 1
	sliceEndingAtIndexOne := names[:1] // gets all the elements ending at index 1

	// Looping through sliceStartingAtIndexOne
	for index, value := range sliceStartingAtIndexOne {
		fmt.Printf("Value using index: %v -- Index: %v -- Value: %v\n", sliceStartingAtIndexOne[index], index, value)
	}

	fmt.Println("")

	for i := 0; i < len(sliceStartingAtIndexOne); i++ {
		fmt.Printf("Index: %v -- Value: %v\n", i, sliceStartingAtIndexOne[i])
	}

	/* ------ */
	fmt.Println("")
	fmt.Println("")
	/* ------ */

	// Looping through sliceEndingAtIndexOne
	for index, value := range sliceEndingAtIndexOne {
		fmt.Printf("Value using index: %v -- Index: %v -- Value: %v\n", sliceEndingAtIndexOne[index], index, value)
	}

	fmt.Println("")

	for i := 0; i < len(sliceEndingAtIndexOne); i++ {
		fmt.Printf("Index: %v -- Value: %v\n", i, sliceEndingAtIndexOne[i])
	}
}