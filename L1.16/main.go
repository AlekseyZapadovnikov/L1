package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)


func main() {
	// тесты
	tests := []struct {
		name string
		in   []int
	}{
		{"empty", []int{}},
		{"one", []int{5}},
		{"sorted", []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}},
		{"dups", []int{2, 3, 2, 3, 2, 3, 2}},
		{"all_equal", []int{1, 1, 1, 1, 1}},
		{"neg_and_big", []int{42, -1, 0, 1000000, -99999}},
		{"odd_len", []int{7, 8, 3, 3, 3, 2, 1, 9, 0, 5, 4, 6}},
		{"even_len", []int{10, 9, 8, 7, 6, 5, 4, 3, 2}},
		{"mixed", []int{12, -3, 45, 0, 23, 23, -3, 88, 7}},
	}

	fmt.Println("=== Deterministic tests ===")
	for _, tt := range tests {
		in := append([]int(nil), tt.in...) // копия input
		want := append([]int(nil), tt.in...)
		sort.Ints(want) // пользуемся готовой сортировкой для проверки

		// запуск quickSort
		if len(in) > 0 {
			quickSortInts(in, 0, len(in)-1)
		} else {
			quickSortInts(in, 0, -1) // безопасный вызов для пустого слайса
		}

		// проверка результата
		ok := reflect.DeepEqual(in, want)
		status := "OK"
		if !ok {
			status = "FAIL"
		}
		fmt.Printf("%-12s %s\n", tt.name, status)

		// для небольших массивов печатаем до/после
		if len(tt.in) <= 12 {
			fmt.Printf("  before: %v\n  after : %v\n  want  : %v\n", tt.in, in, want)
		}
	}

	// Дополнительно: случайный тест побольше
	fmt.Println("\n=== Random test ===")
	rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел
	const N = 1000
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = rand.Intn(2000) - 1000
	}
	want := append([]int(nil), a...)
	sort.Ints(want)

	in := append([]int(nil), a...)
	quickSortInts(in, 0, len(in)-1)

	if reflect.DeepEqual(in, want) {
		fmt.Println("random: OK")
	} else {
		fmt.Println("random: FAIL")
		// печатаем первые 30 элементов для отладки
		fmt.Printf(" first 30 got : %v\n", in[:30])
		fmt.Printf(" first 30 want: %v\n", want[:30])
	}
}


func quickSortInts(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionInts(a, lo, hi)
	quickSortInts(a, lo, p-1)
	quickSortInts(a, p, hi)
}

func partitionInts(a []int, l, r int) int {
	pivot := a[(l+r)/2]
	for l <= r {
		for a[l] < pivot { l++ }
		for a[r] > pivot { r-- }
		if l <= r {
			a[l], a[r] = a[r], a[l]
			l++; r--
		}
	}
	return l
}
