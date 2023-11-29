package main

import (
	"fmt"
	"strings"
)

func main() {
	// var inputStr, regex string

	// fmt.Println("Enter input string")
	// fmt.Scan(&inputStr)

	// fmt.Println("Enter regex string")
	// fmt.Scan(&regex)

	inputStr := "Hello World!, hello go!"
	regex := "ell"

	fmt.Println(strings.ReplaceAll(inputStr, regex, "*"))
}
