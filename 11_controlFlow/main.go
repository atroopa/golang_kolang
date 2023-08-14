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

type MyFloat float64

// functions

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

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

	f := MyFloat(-0.1231)
	fmt.Println(f.Abs())

}
