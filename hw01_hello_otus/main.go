package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	greeting := "Hello, OTUS!"
	reverseGreeting := stringutil.Reverse(greeting)
	fmt.Println(reverseGreeting)
}
