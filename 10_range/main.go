package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for index, value := range pow {
		if index == 7 {
			break
		}

		if index == 3 {
			continue
		}

		fmt.Printf("2^%d = %d \n", index, value)
	}

	Users := map[string]string{
		"name":   "omid",
		"family": "hajavi",
		"id":     "123",
	}

	for index, value := range Users {
		if index == "name" {
			fmt.Println(value)
		}
	}

	cities := map[string]Vertex{}
	cities["Tehran"] = Vertex{53.80, 47.70}

	fmt.Println(cities)

	cities["Tehran"] = Vertex{22.34, 91.01}

	elem, ok := cities["Tehran"]

	if ok {
		fmt.Println(elem)
	}

}

type Vertex struct {
	Latitude, Longitude float64
}
