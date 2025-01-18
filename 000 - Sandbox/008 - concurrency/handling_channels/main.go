package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}

	channels := make([]chan bool, len(numbers))
	errorChannels := make([]chan error, len(numbers))

	for i := range numbers {
		channels[i] = make(chan bool)
		errorChannels[i] = make(chan error)

		go printNumber(numbers[i], channels[i], errorChannels[i])
	}

	for index := range numbers {
		select {
		case err := <- errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- channels[index]:
			fmt.Println("Done!")
		}
	}
}

func printNumber(number int, channel chan bool, errorChannel chan error) {
	if(number == 1) {
		errorChannel <- errors.New("something went wrong")
		return
	}

	fmt.Println(number)
	channel <- true
}