package main

import "fmt"

// create constant variables
const pi = 3.14

const (
	StatusOk      = 200
	StatusCreated = 201
)

// main function
func main() {

	fmt.Printf("this is pi: %f and type is: %T \n", pi, pi)

	// Convert pi to a string
	piStr := fmt.Sprintf("%f", pi)
	fmt.Println("pi as a string:", piStr)
	fmt.Printf("type of pi is : %T \n", piStr)
}

// this functio sum two integer numbers
func sum(a int, b int) int {
	return a + b
}
