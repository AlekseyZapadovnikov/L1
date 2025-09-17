package main

import "fmt"

func main() {
	firstArr := []int{1, 2, 3, 4, 5, 56, 108, 234, 235, 236}
	a := binSearch(4, firstArr)
	b := binSearch(1, firstArr)
	c := binSearch(7, firstArr)
	d := binSearch(236, firstArr)

	fmt.Printf("test1 = %v\n\ntest2 = %v\n\ntest3 = %v\n\ntest4 = %v\n\n", a, b, c, d)
}


// это левый бинпоиск, он возвращает первое число >= данного,
// я просто добавил проверку на точное равенство
// если элемента нет то вернёт -1
func binSearch(num int, arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if num > arr[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	if arr[l] != num {
		return -1
	}
	return l
}