package main

import "fmt"

type interfaceOne interface {
	displayOne()
}

type interfaceTwo interface {
	displayTwo()
}

type interfaceOneAndTwo interface {
	interfaceOne
	interfaceTwo
}

type customer struct {
	firstName string
	lastName  string
}

func (c customer) displayOne() {
	fmt.Printf("displayOne() method: %v - %v", c.firstName, c.lastName);
}

func (c customer) displayTwo() {
	fmt.Printf("displayTwo() method: %v - %v", c.firstName, c.lastName);
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	var i interfaceOneAndTwo
	customer := customer{
		firstName: "Lorem",
		lastName: "Ipsum",
	}

	i = customer

	i.displayOne()
	fmt.Println()
	i.displayTwo()
	fmt.Println()

	var anyExample any

	anyExample = "hello"

	typedValue, ok := anyExample.(string)

	if ok {
		fmt.Printf("%T - %v", typedValue, typedValue)
	}

	fmt.Println()

	integers := add(3, 5)
	floats := add(3.2, 5.4)
	strings := add("3", "5")

	fmt.Println(integers)
	fmt.Println(floats)
	fmt.Println(strings)
}
