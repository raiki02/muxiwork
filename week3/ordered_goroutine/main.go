package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func getOrder(out <-chan int) {
	x, ok := <-out
	if x == 0 || !ok {
		return
	}

	fmt.Printf("Goroutine %d is now working ! \n", x)
	defer wg.Done()
}

func main() {
	order := make(chan int, 20)

	for i := 1; i <= 20; i++ {
		order <- i
	}

	for j := 1; j <= 20; j++ {
		wg.Add(1)
		go getOrder(order)
	}
	close(order)
	wg.Wait()
}
