package main

import "fmt"

func main() {
	fmt.Println("Start of main function")

	defer fmt.Println("This is deferred and will be execute at the very end")

	fmt.Println("Middle of main function")

	defer fmt.Println("This is also deferred and will be execute after the end")

	fmt.Println("End of main function")
}