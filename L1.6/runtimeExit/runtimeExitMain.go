package main

import (
	"runtime"
	"sync"
)
// выход из gorутины через runtime.Goexit() + условие
func main() {
	// Create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if i == 900 {
				println("Goroutine finished")
				runtime.Goexit()
			}
			println("Goroutine is working:", i)
		}
	}()
	wg.Wait()
	println("Main is finished")
}