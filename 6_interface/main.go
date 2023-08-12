package main

import (
	"fmt"
	"time"
)

func timeMap(model interface{}) {
	mapModel, ok := model.(map[string]interface{})
	if ok {
		mapModel["updated_now"] = time.Now()
	}
}

func addToBalance(balance map[string]float64, address string, amount float64) {
	balance[address] += amount
}

func main() {
	foo := map[string]interface{}{
		"Name": "omid",
		"City": "Amesterdom",
	}

	balance := map[string]float64{
		"0x123456789abcdef": 100.0,
		"0x987654321abcdef": 200.0,
	}

	timeMap(foo)
	fmt.Println(foo)

	fmt.Println("Before Adding Balance:")
	fmt.Println(balance)

	addToBalance(balance, "0x123456789abcdef", 50.0)

	fmt.Println("After Adding Balance:")
	fmt.Println(balance)
}
