package main

import "fmt"

func main() {
	sp := []string{"cat", "cat", "dog", "cat", "tree"}
	mp  := make(map[string]struct{})

	for _, v := range sp {
		mp[v] = struct{}{}
	}
	for k := range mp {
		fmt.Println(k)
	}
}