package main

import (
	"fmt"
)

func whriteChan(ch chan <- int) {

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	// close(ch)
}

func main() {
	ch := make(chan int)

	go whriteChan(ch)
	for el := range ch {
		fmt.Println("Прочитали значение из канала:", el)
	}
}
