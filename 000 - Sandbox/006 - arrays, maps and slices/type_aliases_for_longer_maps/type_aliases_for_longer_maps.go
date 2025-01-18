package main

import "fmt"

type myAliasForMap map[string]float64

func (m myAliasForMap) output() {
	fmt.Println(m)
}

func main() {
	myMap := []myAliasForMap{
		{ "Un": 1 },
		{ "Deux": 2 },
		{ "Trois": 3 },
	}

	for index, m := range myMap {
		fmt.Println(index)
		m.output()
	}
}