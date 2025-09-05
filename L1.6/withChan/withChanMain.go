package main

import (
	"sync"
	"time"
)

func main() {
	stopChan := make(chan struct{})
	wg := &sync.WaitGroup{}

	go func(wg *sync.WaitGroup, stopChan <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				return
			default:
				println("working...")
			}
		}
	}(wg, stopChan)

	time.Sleep(2 * time.Second)
	close(stopChan)
}