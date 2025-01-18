package main

import "fmt"

func main() {
	// Creating a map
	languages := map[string]string{
		"French":     "Bonjour!",
		"English":    "Hello!",
	}

	// Adding an item
	languages["Portuguese"] = "Oi!"

	// Looping through a map
	for key, value := range languages {
		fmt.Println(key, value)
	}

	// Deleting an item
	delete(languages, "English")
	fmt.Println("")

	for key, value := range languages {
		fmt.Println(key, value)
	}
}