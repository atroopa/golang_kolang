package main

import (
	"fmt"
)

// ========types===================
type User struct {
	Name   string
	Family string
	age    int64
}

type Player struct {
	User
	PlayerID int64
}

// =========main====================

func main() {

	omid := &Player{}
	omid.Name = "omid"
	omid.PrintName()

	var a [3]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "!!"
	fmt.Println(a)

	b := [3]string{"hello", "World", "!!"}
	fmt.Println(b)

	c := [2][3]string{
		{"this", "is", "first"},
		{"this", "is", "second"},
	}
	fmt.Println(c[0])
	fmt.Println(c[1])

}

// ========functions=================

func (u *Player) PrintName() {
	fmt.Println(u.Name)
}
