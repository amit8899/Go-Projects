package main

import "fmt"

func FibonacciSeries(ch chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	num := 10

	ch := make(chan int)
	quit := make(chan int)

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println(<-ch) // read from channel ch
		}
		quit <- 1
	}(num)

	FibonacciSeries(ch, quit)
}
