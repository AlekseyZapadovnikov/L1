package main

import (
	"fmt"
)

// xor swap
func main() {
	a := 1
	b := 2
	a = a ^ b
	b = a ^ b
	a = a ^ b

	// а вообще можно просто a, b = b, a
	fmt.Println(a, b)
}