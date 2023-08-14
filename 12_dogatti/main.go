package main

import (
	"fmt"
)

// types ===============================
type User struct {
	Name   string
	Family string
	Age    int
}

type Namer interface {
	FullName() string
}

// functions ==================================
func (u *User) FullName() string {
	return fmt.Sprintf("Dear %s %s", u.Name, u.Family)
}

func Greeting(namer Namer) string {
	return fmt.Sprintf("Greeting %s", namer.FullName())
}

// main  =====================================
func main() {

	user := &User{"omid", "hajavi", 30}
	fmt.Println(user.FullName())

	fmt.Println(Greeting(user))

}
