package main

import (
	"fmt"
)

func main() {
	// set limit to maksimum channel receive amount of data
	limit := 5
	// create a channel to receive value of odd number
	// this channel would filled up with every value that sent by func oddNumber
	// channel also make process go routine is synchronous
	// so the main go routine is waiting/blocked till channel in this go routine is filled up
	o := make(chan int)
	// create a channel to receive value of even number
	// this channel would filled up with every value that sent by func evenNumber
	// channel also make process go routine is synchronous
	// so the main go routine is waiting/blocked till channel in this go routine is filled up
	e := make(chan int)

	// spawn go routine to asynchronouze oddNumber
	// but channel o will sync til filled up
	go oddNumber(o)
	// spawn go routine to asynchronouze evenNumber
	// but channel e will sync til filled up
	go evenNumber(e)
	// looping variable to receive data from channel
	for i := 0; i < limit; i++ {
		odd := <-o
		even := <-e
		fmt.Printf("Odd: %d\n", odd)
		fmt.Printf("Even: %d\n", even)
	}

	// this go routine is an example of without using channel
	// this would make a process of main go routine won't wait this go routine
	// the main go routine doesn't care wheter this proccess is done or not.
	go printHello()
}

// oddNumber is a function that works to send value to channel o
// every value that this function have would immediately send to channel o
func oddNumber(o chan int) {
	// looping forever to find number that needed
	// if the number is an odd then send number to channel
	for i := 0; ; i++ {
		if i%2 == 1 {
			o <- i
		}
	}
}

// evenNumber is a function that works to send value to channel e
// every value that this function have would immediately send to channel e
func evenNumber(e chan int) {
	// looping forever to find number that needed
	// if the number is an even then send number to channel
	for i := 0; ; i++ {
		if i%2 == 0 {
			e <- i
		}
	}
}

func printHello() {
	fmt.Println("Hello World")
}
