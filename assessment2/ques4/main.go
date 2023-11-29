package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go numbers()
	go alphabets()
	wg.Wait()
}

func numbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func alphabets() {
	chars := []rune{'a', 'b', 'c', 'd', 'e'}
	for _, v := range chars {
		fmt.Printf("%c\n", v)
		time.Sleep(80 * time.Millisecond)
	}

	wg.Done()
}
