package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	tmpArr := []rune(str)
	// выводим в обратном порядке
	for i := len(tmpArr) - 1; i >= 0; i-- {
		fmt.Print(string(tmpArr[i]))
	}
}

/* будет долго работать на большом наборе данных,
тк fmt считывает и выводит по 1 байду (последовательно, а не  буфером)*/