package main

import (
	"fmt"
	"sync"
)

func main() {
	sp := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	wg.Add(len(sp)) // чтобы не вызывать функкцию несколько раз
	for _, val := range sp {
		go func(v int) {
			defer wg.Done() // откладываем вызов, чтобы -1 произошло, когда мы завершим горутину
			fmt.Printf("было число %v теперь это число в квадрате = %v \n", v, v*v)
		}(val)
	}
	wg.Wait()
}