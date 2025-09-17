package main

import "fmt"

// Удаление с сохранением порядка
func DeleteAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s // или panic
	}
	// сдвигаем хвост влево
	copy(s[i:], s[i+1:])
	// Обнуление удаляет ссылку (важно для ссылочных типов)
	// и даёт возможность сборщику мусора освободить память.
	var zero T
	s[len(s)-1] = zero
	// укоротим слайс на 1
	return s[:len(s)-1]
}

// Удаление без сохранения порядка (быстрее O(1), но портит порядок)
func DeleteAtUnordered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	// заменяем i-ый элемент последним
	s[i] = s[len(s)-1]
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

type Node struct {
	Name string
}

func main() {
	// Пример с сохранением порядка
	nodes := []*Node{{"a"}, {"b"}, {"c"}, {"d"}}
	fmt.Println("before:", nodes)
	nodes = DeleteAt(nodes, 1) // удаляем "b"
	fmt.Println("after (preserve order):", nodes)

	// Пример без сохранения порядка
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("nums before:", nums)
	nums = DeleteAtUnordered(nums, 1) // удаляем элемент с индексом 1
	fmt.Println("nums after (unordered):", nums)
}
