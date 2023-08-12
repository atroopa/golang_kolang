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
	example = "kirkos"
	fmt.Println(example)

}

// functions
func ChangeFamily(familyAddress *string, newFamily string) {
	*familyAddress = newFamily
}
