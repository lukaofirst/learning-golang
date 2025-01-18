package main

import "fmt"

type transformFn func(int) int

func main() {
	// Function as values and function types
	fmt.Println(principalFunction(100, anotherFunction))
	
	// Return functions as values
	anotherFunction := functionThatReturnsAnotherFunction()
	fmt.Println(anotherFunction(3))

	// Anonymous function
	anonymousFunctionResult := anonymousFunction(10, func (number int) int {
		return number * 3
	})

	fmt.Println(anonymousFunctionResult)

	// Variadic function
	fmt.Println(variadicFunctionToSumUp(120, 340, 440, 1))

	// Splitting slices into parameters values
	numbers := []int{1, 3, 5}
	fmt.Println(pullingOutFromSliceToSumUp(numbers...))
}

/* --- Function as values and function types --- */
func principalFunction(number int, transform transformFn) int {
	return transform(number)
}

func anotherFunction(numberToTransform int) int {
	return numberToTransform + 1
}

/* --- Return functions as values --- */
func functionThatReturnsAnotherFunction() func(int) int {
	return anotherFunction
}

/* --- Anonymous function --- */
func anonymousFunction(number2 int, fn func (number int) int) int {
	return fn(number2)
}

/* --- Variadic function --- */
func variadicFunctionToSumUp(numbers ...int) int {
	sum := 0
	
	for _, val := range numbers {
		sum += val
	}
	
	return sum
}

/* --- Splitting slices into parameters values --- */
func pullingOutFromSliceToSumUp(numbers ...int) int {
	sum := 0
	
	for _, val := range numbers {
		sum += val
	}
	
	return sum
}