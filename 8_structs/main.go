package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// ===================== types =============================
type NewStruct struct {
	Name   string   `json:"name"`
	Family string   `json:"family"`
	Model  newModel `json:"model"`
}

type Stringer interface {
}

type newModel struct {
	ID        uint64
	CreatedAt time.Time
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

	// Print JSON representation of NewStruct
	jsonOutput, err := newStruct.ToJSON()
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
	} else {
		fmt.Println("====== JSON output ==========")
		fmt.Println(jsonOutput)
	}

}

// =============== functions ================================
func (ns *NewStruct) SayHello(newName string, newFamily string) string {
	ns.Name = newName
	ns.Family = newFamily
	rand.Seed(time.Now().UnixNano())
	ns.Model.ID = rand.Uint64()
	ns.Model.CreatedAt = time.Now()
	result := fmt.Sprintf("Hello Name: %s Family: %s id: %d time: %s", ns.Name, ns.Family, ns.Model.ID, ns.Model.CreatedAt)
	return result
}

// ToJSON returns the JSON representation of the NewStruct instance
func (ns *NewStruct) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(ns)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
