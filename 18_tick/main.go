package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello from tick")

	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BooM!")
			return
		default:
			fmt.Println("----------.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
