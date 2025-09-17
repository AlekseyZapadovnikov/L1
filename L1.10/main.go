package main

import (
	"fmt"
	"sort"
)

func main() {
	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // данный массив

	groups := make(map[int][]float64) // это будет ответ
	keys := make([]int, 0, len(seq))  // ключи для сортировки

	// группируем числа по ключу
	for _, v := range seq {
		key := giveGroup(v)
		groups[key] = append(groups[key], v)
	}

	// сортируем ключи для дальнейшей сортировки
	for k := range groups {
		keys = append(keys, k)
	}

	// сортируем ключи
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}

// функция giveGroup() определяет группу для каждого числа
func giveGroup(in float64) int {
	if in >= 0.0 {
		val := int(in)
		val = val - val%10
		return val
	} else {
		in *= -1
		val := int(in)
		val = val - val%10
		return val * -1
	}

}

/* асимптотическая сложность O(n * log(n)) из за сортировки ключей */