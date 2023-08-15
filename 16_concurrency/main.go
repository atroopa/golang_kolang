package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

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

	c := make(chan int, 10)
	fibonacci(cap(c), c)

	fmt.Println("Before Execute Printing")
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("After Execute Printing")

}
