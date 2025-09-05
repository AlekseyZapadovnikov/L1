package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	sec := flag.Int("sec", 3, "параметр N - количество секунд до завершения работы")
	flag.Parse()

	var wg sync.WaitGroup
	ch := make(chan int, 100)

	// Используем time.After для таймаута
	timeout := time.After(time.Duration(*sec) * time.Second)

	wg.Add(1)
	go writeGoroutine(ch, timeout, &wg)

	wg.Add(1)
	go readGoroutine(ch, &wg)

	wg.Wait()
	fmt.Println("done")
}

func writeGoroutine(ch chan<- int, timeout <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // writer закрывает канал при выходе

	i := 0
	for {
		select {
		case <-timeout:
			// время вышло — завершаемся и закрываем канал (через defer)
			return
		case ch <- i + 1:
			i++
		}
	}
}

func readGoroutine(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
