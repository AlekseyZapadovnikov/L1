package main

import (
	"fmt"
	"os"
	"strconv"
)

// пишет в канал
func worker(ch <-chan int, id int) {
	for val := range ch {
		fmt.Printf("worker %v прочитал значение %v \n", id, val)
	}
}

func main() {
	ch := make(chan int)
	workerAmount, err := strconv.Atoi(os.Args[len(os.Args)-1]) // преобразуем аргумент командной строки в число

	if err != nil || workerAmount <= 0 {
		fmt.Println("вы ввели невалидное число число")
		return
	}

	for i := 1; i <= workerAmount; i++ {
		go worker(ch, i)
	}

	counter := 0
	for {
		counter++
		ch <- counter
	}
}
