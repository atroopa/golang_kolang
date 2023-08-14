package main

import (
	"fmt"
)

// types =============================================================
type User struct {
	Name   string
	Family string
	Age    int
	ID     int
}

type Namer interface {
	FullName() string
}

type Player struct {
	User
	PlayerID int
}

type Identity interface {
	GetID() string
}

// functions =======================================================
func (u *User) FullName() string {
	return fmt.Sprintf("Dear %s %s \n", u.Name, u.Family)
}

func (u *User) GetID() string {
	return fmt.Sprintf("Dear %d \n", u.ID)
}

func (u *Player) GetID() string {
	return fmt.Sprintf("Dear %d \n", u.PlayerID)
}

func ShowID(indentity Identity) string {
	return fmt.Sprintf("ID %s", indentity.GetID())
}

// main  ===========================================================
func main() {

	user := &User{"omid", "hajavi", 30, 1003}
	fmt.Println(ShowID(user))

}
