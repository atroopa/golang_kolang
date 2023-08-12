package main

import "fmt"

var (
	family  string = "khajavi"
	example string = "address"
)

// main function
func main() {

	ChangeFamily(&family, "khajevand")
	fmt.Println(family)
	fmt.Printf("this is address of example: %p \n", &example)
	fmt.Println("==============================")

	var num int = 42
	var ptr *int

	ptr = &num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value of ptr: %p\n", ptr)
	fmt.Printf("Value pointed by ptr: %d\n", *ptr)
}

// functions
func ChangeFamily(familyAddress *string, newFamily string) {
	*familyAddress = newFamily
}
