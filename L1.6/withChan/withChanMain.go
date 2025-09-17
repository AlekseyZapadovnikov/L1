package main

import (
	"time"
)

func main() {
	stopChan := make(chan struct{})

	go func(stopChan <-chan struct{}) {
		for {
			select {
			case <-stopChan:
				return
			default:
				println("working...")
			}
		}
	}(stopChan)

	time.Sleep(2 * time.Second)
	close(stopChan)
}