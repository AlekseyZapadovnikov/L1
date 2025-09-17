package main

import (
	"fmt"
	"sync"
	"time"
)

// выход из goroutine по условию time.Now().Second()%7 == 0

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// запускаем горутину
	go func() {
		defer wg.Done()
		// бесконечный цикл
		for {
			// если выполняется условие то выходим из горутины
			if time.Now().Second()%7 == 0 {
				fmt.Printf("we are out from goroutine, current second: %d\n", time.Now().Second())
				return
			}
			fmt.Printf("we are in goroutine, current second: %d\n", time.Now().Second())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}