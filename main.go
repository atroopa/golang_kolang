package main

import (
	"doost/packages/test"
	"doost/packages/testPackage"
	"fmt"
)

func main() {

	fmt.Println(testPackage.Greet())
	fmt.Println(test.SayHello())
	fmt.Println(test.Pi)

	name := "vali"
	message, err := test.SeyHelloWithName(name)

	// check for error
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print result if no has error
	fmt.Println(message)
}
