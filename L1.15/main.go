// // неправильный вариант

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func createHugeString(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = 'A'
// 	}
// 	return string(b)
// }

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// 	fmt.Printf("%v\n", reflect.TypeOf(justString))
// }


/* проблема в том, что при v[:100] не создаётся копия строки, 
а создаётся срез - структура которая содержит (указатель на исходную строку и длину)
то есть при изменении исходной строки, изменится и срез
это "во-первых", во-вторых, мы не освобождаем память
я имею в виду, что исходная строка остаётся в памяти, хотя исходя из логики она нам не нужна
нам нужны только первые 100 её символов
Более того s[:100] скопирует 100 байт, а не символов (1 байт != 1 символу)
в UNICODE символы весят от 1 до 4 байт
вот 3 ключевые проблемы
*/

package main

import "fmt"

var justString string

func createHugeString(n int) string {
	// пример: большая строка
	b := make([]rune, n)
	for i := range b {
		b[i] = '世' // 3 байта в UTF-8
	}
	return string(b)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// скопируем только первые 100 символов в маленький буфер
	want := 100
	b := make([]rune, want)
	for i := 0; i < want; i++ {
		b[i] = rune(v[i])
	}
	justString = string(b) // теперь justString не держит большой буфер
}

func main() {
	someFunc()
	fmt.Println(len(justString), "bytes")
}