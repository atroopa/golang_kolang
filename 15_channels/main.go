package main

import (
	"fmt"
	"time"
)

func SendValueToChannel(channel chan int, value int) {

	for i := 0; i < 10; i++ {
		channel <- value * i
		time.Sleep(time.Second)
	}
}

func ReceiveFromChannel(channel chan int) {
	for {
		var value int = <-channel
		fmt.Println(value)
	}

}

func main() {

	channel1 := make(chan int, 100)
	var value1 int = 1000

	go SendValueToChannel(channel1, value1)
	go ReceiveFromChannel(channel1)

}
