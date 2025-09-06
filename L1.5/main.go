package main

import (
	"context"
	"flag"
	"sync"
	"time"
)

var timer int

func main() {
	flag.IntVar(&timer, "sec", 3, "параметр N - количество секунд до завершения работы") // флаг при запуске --timer
	flag.Parse()

	wg := &sync.WaitGroup{}
	ch := make(chan int, 100)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timer)) // создаём контекст с таймаутом
	defer cancel() // если main завершится раньше, то отменим контекст (частая практика)
	
	// запускаем горутины
	wg.Add(2)
	go writeGorutine(ch, ctx, wg)
	go readGorutine(ch, wg)
	wg.Wait()
}

// пишет в канал
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

// читает из канала
func readGorutine(ch <-chan int, wg *sync.WaitGroup) {
	for el := range ch {
		println(el)
	}
	wg.Done()
}
