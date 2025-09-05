package main

import (
	"context"
	"flag"
	"sync"
	"time"
)

var timer int

func main() {
	flag.IntVar(&timer, "sec", 3, "параметр N - количество секунд до завершения работы")
	flag.Parse()

	wg := &sync.WaitGroup{}
	ch := make(chan int, 100)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timer))
	defer cancel()
	wg.Add(1)
	go writeGorutine(ch, ctx, wg)
	wg.Add(1)
	go readGorutine(ch, wg)
	wg.Wait()
}

func writeGorutine(ch chan<- int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	c := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			c++
			ch <- c
		}
	}
}

func readGorutine(ch <-chan int, wg *sync.WaitGroup) {
	for el := range ch {
		println(el)
	}
	wg.Done()
}
