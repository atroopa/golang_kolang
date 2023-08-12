package test

import "fmt"

var (
	Pi     float64 = 3.14
	result string
)

func SayHello() string {
	return "Hello from test package!"
}

func SeyHelloWithName(name string) (string, error) {
	result = fmt.Sprintf("Fuck You %s", name)
	return result, nil
}
