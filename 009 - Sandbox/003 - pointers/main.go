package main

import "fmt"

func main() {
	age := 25

	printPointersAndValues(&age)
}

func printPointersAndValues(age *int) {
	agePointer := age

	fmt.Println(agePointer)
	fmt.Println(*age)
}

// age := 25
// var agePointer *int
// agePointer = &age

// fmt.Println(agePointer)
// fmt.Println(age)