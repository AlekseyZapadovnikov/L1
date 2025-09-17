package main

import "fmt"

/* будем использовать некий set (в го он реализован через map) */

func main() {
	sp := []string{"cat", "cat", "dog", "cat", "tree"} // исходный массив
	mp  := make(map[string]struct{})

	// оставляем только уникальные слова
	for _, v := range sp {
		mp[v] = struct{}{}
	}
	// результирующее множество
	ans := make([]string, 0, 4)
	for k := range mp {
		ans = append(ans, k)
	}
	fmt.Println(ans)
}