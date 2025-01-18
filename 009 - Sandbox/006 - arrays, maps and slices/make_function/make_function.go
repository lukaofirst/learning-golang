package main

import "fmt"

func main() {
	// Creating an array using make function
	programmingLangsInitArray := make([]string, 0, 5)

	// Doing like this you'll receive an out bound error,
	// because you're creating an array with capacity of 5, but with length of 0
	// resulting in a not accessible array

	// programmingLangsInitArray[0] = "C#"
	// programmingLangsInitArray[1] = "JavaScript"
	// programmingLangsInitArray[2] = "Go"

	// but you can do it using the append function
	programmingLangs := append(programmingLangsInitArray, "C#", "JavaScript", "Go")

	for i, v := range programmingLangs {
		fmt.Println(i, v)
	}

	fmt.Println("")

	// Creating a map using make function
	numbersInFrenchMap := make(map[int64]string)
	numbersInFrenchMap[0] = "Un"
	numbersInFrenchMap[1] = "Deux"
	numbersInFrenchMap[2] = "Trois"
	numbersInFrenchMap[3] = "Quatre"
	numbersInFrenchMap[4] = "Cinq"

	// Unless you make a loop in the map, you cannot get directly a slice from a map

	for i, v := range numbersInFrenchMap {
		fmt.Println(i, v)
	}
}