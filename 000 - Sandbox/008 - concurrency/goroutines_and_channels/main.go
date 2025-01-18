package main

import (
	"fmt"
	"time"
)

func main() {
	doneChannel := make([]chan bool, 4)

	doneChannel[0] = make(chan bool)
	go printMessage("hello 1", doneChannel[0])
	doneChannel[1] = make(chan bool)
	go printMessage("hello 2", doneChannel[1])
	doneChannel[2] = make(chan bool)
	go printMessageTakesLonger("hello ... I need more time", doneChannel[2])
	doneChannel[3] = make(chan bool)
	go printMessage("hello 3", doneChannel[3])

	for _, done := range doneChannel {
		<- done
	}
}

func printMessage(message string, channel chan bool) {
	fmt.Println(message)
	channel <- true
}

func printMessageTakesLonger(message string, channel chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println(message)
	channel <- true
}