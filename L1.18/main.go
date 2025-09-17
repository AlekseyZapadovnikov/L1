package main

import (
	"sync"
	"sync/atomic"
)

func myInc(counter *concurentCounter, wg *sync.WaitGroup) {
	counter.inc()
	wg.Done()
}

// я написал обёртку над atomic.Int64
// тк все операции атомарные, то не может произойти гонки
func main() {
	counter := concurentCounter{counter: atomic.Int64{}}
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go myInc(&counter, &wg)
	}

	wg.Wait()
	println(counter.counter.Load())
}


type concurentCounter struct {
	counter atomic.Int64
}

func (concCounter *concurentCounter) inc() {
	concCounter.counter.Add(1)
}