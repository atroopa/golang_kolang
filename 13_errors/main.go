package main

import (
	"fmt"
	"time"
)

// type err interface {
// 	Error() string
// }

type MyError struct {
	CreatedAt time.Time
	Body      string
}

// functions =====================
func (me *MyError) Error() string {
	return fmt.Sprintf("At %s Error: %s", me.CreatedAt, me.Body)
}

func RunMyApp() error {
	newError := &MyError{time.Now(), "we Have funny Error Here !"}
	return newError
}

// main ===========================
func main() {

	myApp := RunMyApp()
	fmt.Println(myApp)

}
