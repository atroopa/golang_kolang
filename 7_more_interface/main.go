package main

import (
	"fmt"
)

type savedStr interface {
	Test() string
}

type fakeString struct {
	content string
}

func (small *fakeString) Test() string {
	return small.content
}

func printSaving(value interface{}) {
	result, ok := value.(savedStr)

	if ok {
		fmt.Println(result.Test())
	}
}

func main() {
	sample := &fakeString{
		content: "this is sample text",
	}

	printSaving(sample)
}
