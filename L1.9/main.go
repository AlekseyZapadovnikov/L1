package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	
	wg.Add(3)
	go goWriter(arr, ch1, &wg)
	go goSquare(ch1, ch2, &wg)
	go goPrinter(ch2, &wg)

	wg.Wait()
}

func goWriter(arr []int, ch chan<- int, wg *sync.WaitGroup) {
	for _, v := range arr {
		ch <- v
	}
	close(ch)
	wg.Done()
}

func goSquare(ch <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for v := range ch {
		out <- v * v
	}
	close(out)
	wg.Done()
}

func goPrinter(out <-chan int, wg *sync.WaitGroup) {
	for v := range out {
		fmt.Println(v)
	}
	wg.Done()
}
