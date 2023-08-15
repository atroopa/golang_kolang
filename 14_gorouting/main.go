package main

import (
	"fmt"
	"time"
)

func PrintMyText(text string) {
	for i := 0; i < len(text); i++ {
		fmt.Println(string(text[i]))
		time.Sleep(time.Second)
	}
}

func main() {

	PrintMyText("hello")
	PrintMyText("world")
}
