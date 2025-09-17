package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // исходный массив
	// канал для передачи чисел
	ch1 := make(chan int)
	ch2 := make(chan int)
	// WaitGroup для ожидания завершения всех горутин
	wg := sync.WaitGroup{}
	
	wg.Add(3)
	go goWriter(arr, ch1, &wg)
	go goSquare(ch1, ch2, &wg)
	go goPrinter(ch2, &wg)

	// ждём завершения всех горутин
	// нужно для корректного чтения и записи из канала
	wg.Wait()
}

// функция, которая пишет числа из массива в канал
func goWriter(arr []int, ch chan<- int, wg *sync.WaitGroup) {
	for _, v := range arr {
		ch <- v
	}
	close(ch)
	wg.Done()
}

// функция, которая читает числа из канала, возводит их в квадрат и пишет в другой канал
func goSquare(ch <-chan int, out chan<- int, wg *sync.WaitGroup) {
	for v := range ch {
		out <- v * v
	}
	close(out)
	wg.Done()
}

// функция, которая читает числа из канала и печатает их
func goPrinter(out <-chan int, wg *sync.WaitGroup) {
	for v := range out {
		fmt.Println(v)
	}
	wg.Done()
}
