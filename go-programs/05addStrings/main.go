package main

import "fmt"

func main() {
	var firstStr, secStr string
	fmt.Println("Enter first string")
	fmt.Scanln(&firstStr)

	fmt.Println("Enter second string")
	fmt.Scanln(&secStr)

	// both are good
	fmt.Print("after adding: ", firstStr+secStr)
	fmt.Print("after adding: ", firstStr, secStr)
}
