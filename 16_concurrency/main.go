package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from concurrency")

	channel := make(chan int, 1)

	go func() {
		channel <- 100

	}()

	value := <-channel
	fmt.Println(value)

	channel <- 200
	malue := <-channel
	fmt.Println(malue)
}
