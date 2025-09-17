package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // разделять на слова (пробелы, переводы строки и т.д.)

	words := make([]string, 0, 8)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return
	}

	// выводим в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
}
