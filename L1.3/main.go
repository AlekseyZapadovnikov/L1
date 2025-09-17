package main

import (
	"fmt"
	"os"
	"strconv"
)

// читает из канала и выводит какой воркер что прочитал
func worker(ch <-chan int, id int) {
	for val := range ch {
		fmt.Printf("worker %v прочитал значение %v \n", id, val)
	}
}

func main() {
	ch := make(chan int)
	workerAmount, err := strconv.Atoi(os.Args[len(os.Args)-1]) // преобразуем аргумент командной строки в число

	// валидация числа
	if err != nil || workerAmount <= 0 {
		fmt.Println("вы ввели невалидное число число")
		return
	}

	// запускаем n воркеров
	for i := 1; i <= workerAmount; i++ {
		go worker(ch, i)
	}

	counter := 0
	// бесконечный цикл, который пишет в канал
	for {
		counter++
		ch <- counter
	}
}
