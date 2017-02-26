package main

import (
	"log"
	"runtime"
	"sync"
)

var MAXGOROUTINES int = 5

func myFunc(val int, wg *sync.WaitGroup) {
	log.Printf("Value %d", val)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Create a channel of capacity 5
	tokens := make(chan bool, MAXGOROUTINES)

	// slice of 10 items
	aSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range aSlice {
		wg.Add(1)
		log.Printf("Number of goroutines: %d", runtime.NumGoroutine())
		// Get a token
		tokens <- true
		go func(val int) {
			myFunc(val, &wg)
			// release the token
			<-tokens
		}(val)
	}
}
