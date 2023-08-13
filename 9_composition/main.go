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
	omid := new(Player)
	omid.Name = "omid"
	omid.PrintName()
}

// ========functions=================

func (u *Player) PrintName() {
	fmt.Println(u.Name)
}
