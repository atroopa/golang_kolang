package main

import "fmt"

var i = 42
var f = float64(i)
var u = uint(f)

func main() {
	const format = "%T(%v)\n"

	fmt.Printf(format, i, i)
	fmt.Printf(format, f, f)
	fmt.Printf(format, u, u)
}
