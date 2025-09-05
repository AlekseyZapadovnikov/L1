package main

import (
	"context"
	"time"
	"sync"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()
	wg := &sync.WaitGroup{}

	wg.Add(1)
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

	wg.Wait()
}