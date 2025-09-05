package main

import (
	"fmt"
	"sync"
)

func main() {
	sp := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, val := range sp {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Printf("было число %v теперь это число в квадрате = %v \n", v, v*v)
		}(val)
	}
	wg.Wait()
}