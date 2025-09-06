package main

import (
	"fmt"
	"sort"
)

func main() {
	seq := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	keys := make([]int, 0, len(seq))

	for _, v := range seq {
		key := giveGroup(v)
		groups[key] = append(groups[key], v)
	}

	for k := range groups {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}

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
