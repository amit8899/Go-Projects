package main

import "fmt"

var factVal uint64 = 1 // uint64 is the set of all unsigned 64-bit integers.
// Range: 0 through 18446744073709551615.

var n int

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}

func factorialbyrecursion(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorialbyrecursion(i-1)
}

func main() {
	fmt.Print("Enter a positive integer between 0 - 50 : ")
	fmt.Scan(&n)
	fmt.Println("Factorial is: ", factorial(n))

	fmt.Printf("Factorial of %d by recursion is: %d", n, factorialbyrecursion(n))
}
