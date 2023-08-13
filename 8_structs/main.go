package main

import (
	"fmt"
)

// ===================== types =============================
type NewStruct struct {
	Name   string
	Family string
}

type Stringer interface {
}

// =============== main function ==========================
func main() {
	fmt.Print("====== start ===== \n")
	newStruct := &NewStruct{}
	printer := newStruct.SayHello("Golnaz", "Arbabi")
	fmt.Println(printer)
	fmt.Println("====== variables ==========")
	fmt.Println(newStruct.Name)
	fmt.Println(newStruct.Family)
}

// =============== functions ================================
func (ns *NewStruct) SayHello(newName string, newFamily string) string {
	ns.Name = newName
	ns.Family = newFamily
	result := fmt.Sprintf("Hello Name: %s Family: %s", ns.Name, ns.Family)
	return result
}
