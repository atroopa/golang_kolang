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

func main() {

	foo := map[string]interface{}{
		"Name": "omid",
		"City": "Amesterdom",
	}
	timeMap(foo)

	fmt.Println(foo)

}
