package main

// imports
import (
	"fmt"
)

// structs for make subjects
type Artist struct {
	Name, Genre string
	Songs       int
}

// function added a step to aongs
func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

// main function
func main() {
	me := &Artist{Name: "Matt", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %d th songs \n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs \n", me.Name, me.Songs)
	fmt.Println(&me)
}
