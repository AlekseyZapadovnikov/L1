package main

import (
	"context"
	"time"
	"sync"
)

// остановка горутины через контекст с дедлайном
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	// запускаем горутину, которая будет работать, пока контекст не будет отменён
	// а отменён он будет по прошествии 3 секунд
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
				case <-ctx.Done():
					println("Goroutine finished")
					return
				default:
					println("Goroutine is working...")
					time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	wg.Wait() // нужно чтобы горутина отработала
}