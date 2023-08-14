package main

// imports
import (
	"fmt"
	"strings"
)

// types
type User struct {
	FirstName string
	LastName  string
}

type MyString string

// functions
func (s MyString) Uppercase() string {
	return strings.ToUpper(string(s))
}

func Test() string {
	return "test"
}

func (u User) Greeting() string {
	return fmt.Sprintf("Salam %s %s", u.FirstName, u.LastName)
}

// main
func main() {
	user := User{
		FirstName: "Omid",
		LastName:  "Hajavi",
	}

	fmt.Println(user.Greeting())
	fmt.Println(Test())

	myString := MyString("salam")
	fmt.Println(myString.Uppercase())

}
