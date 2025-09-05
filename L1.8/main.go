package main

import (
	"fmt"
)

func main() {
	var n, bit, oneOrZero, mask, ans int64
	fmt.Println("введите исходное число")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	fmt.Println("введите бит, который нужно поменять (0-63)")
	_, err = fmt.Scan(&bit)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	fmt.Printf("введите на какой бит вы хотите поменть %v бит (0 или 1)\n", bit)
	_, err = fmt.Scan(&oneOrZero)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	mask = 1
	if oneOrZero == 0 {
		mask = 1
		r := 63 - int(bit)
		for i := 0; i < r; i++ {
			mask = (mask<<1) + 1
		}
		mask = mask << 1

		for i := 0; i < int(bit); i++ {
			mask = (mask << 1) + 1
		}
		ans = n & int64(mask)
	} else {
		for i := 0; i < int(bit); i++ {
			mask = (mask<<1)
		}
		ans = n | mask
	}

	fmt.Println("Ответ:", ans)
}