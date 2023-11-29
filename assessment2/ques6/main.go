package main

import (
	"fmt"
	"reflect"
)

func main() {
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{1, 2, 3, 4, 5}

	s1 := swaptwoelements(data1) //1

	s2 := reverseSlice(data2) //2

	if reflect.DeepEqual(s1, s2) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}
}

func reverseSlice(data []int) []int {
	fmt.Printf("Original Slice: %v\n", data)

	swapF := reflect.Swapper(data)

	for i := 0; i < len(data)/2; i++ {
		swapF(i, len(data)-1-i)
	}

	// printing the values
	fmt.Printf("Reversed Slice: %v\n", data)

	return data
}

func swaptwoelements(data []int) []int {
	fmt.Printf("Before swap: %v\n", data)
	swapF := reflect.Swapper(data)

	swapF(2, 4)

	fmt.Printf("After swap: %v\n", data)

	return data
}
