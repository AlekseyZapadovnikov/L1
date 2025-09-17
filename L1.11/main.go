package main

import "fmt"

/* пересечение множеств
   сначала создаём map на основе второго массива, чтобы узнавать о наличии элемента за O(1)
   потом проходим по первому массиву и проверяем наличие элемента во втором
*/

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	set := make(map[int]struct{})

	for _, v := range b {
		set[v] = struct{}{}
	}

	ans := make([]int, 0, min(len(b), len(a)))
	for _, v := range a {
		if _, ok := set[v]; ok {
			ans = append(ans, v)
		}
	}
	fmt.Printf("%v\n", ans)
}

// асимптотика O(n),
// а если делать for {for {}}, 
// то будет O(n*m) то есть квадрат - это оч плохо