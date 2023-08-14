package main

import (
	"fmt"
)

// types ===============================
type User struct {
	FirstName, LastName string
}

// functions ==================================
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

// main  =====================================
func main() {

	u := &User{"omid", "hajavi"}
	fmt.Println(u.Greeting())

}
