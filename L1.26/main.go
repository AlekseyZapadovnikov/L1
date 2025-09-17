package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
дефолтная зажача на set, в Go set реализован с помощью map[string]struct{}
бежим по строке и добавляем буквы в множество, если в множестве уже есть такая буква
то есть в строке ранее мы уже встречали данную букву тогда таких букв (одинаковых) хотябы 2
сразу же выводим false и завершаем выполнение программы
*/

func main() {
	mp := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var letter string
		fmt.Sscanf(scanner.Text(), "%c", &letter)
		letter = strings.ToLower(letter)
		if _, exists := mp[letter]; exists {
			fmt.Println("false")
			return
		}
		mp[letter] = struct{}{}
	}
	fmt.Println("true")
}
