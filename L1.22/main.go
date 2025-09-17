package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin) // читаем из StdIn (возможно тут bufioReader не факт, что оправдан, дольше памть под него выделять + там целый буфер)

	// ---- инпут ----
	fmt.Print("Введите a: ")
	var sa string
	if _, err := fmt.Fscan(in, &sa); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения a:", err)
		return
	}

	fmt.Print("Введите b: ")
	var sb string
	if _, err := fmt.Fscan(in, &sb); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения b:", err)
		return
	}

	a := new(big.Int)
	b := new(big.Int) // а и b - это указатели на big.Int

	// SetString парсит строку в big.Int, вторым аргументом идёт основание системы счисления (10 - десятичная)
	// SetString возвращает указатель на big.Int и bool (успешно ли прошло преобразование)
	// если неуспешно, то big.Int остаётся равен 0

	if _, ok := a.SetString(sa, 10); !ok {
		fmt.Fprintln(os.Stderr, "не удалось распарсить a")
		return
	}
	if _, ok := b.SetString(sb, 10); !ok {
		fmt.Fprintln(os.Stderr, "не удалось распарсить b")
		return
	}

	// вычисляем сумму, разность и произведение
	// это тесты

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)

	fmt.Println("a =", a.String())
	fmt.Println("b =", b.String())
	fmt.Println()
	fmt.Println("a + b =", sum.String())
	fmt.Println("a - b =", diff.String())
	fmt.Println("a * b =", prod.String())

	if b.Sign() == 0 {
		fmt.Println("Деление на ноль: частное и остаток не вычисляются.")
	} else {
		quo := new(big.Int).Div(a, b)  // целочисленное деление (частное)
		rem := new(big.Int).Mod(a, b)  // остаток (неотрицательный по модулю б)
		fmt.Println("a / b (целочисленное) =", quo.String())
		fmt.Println("a % b =", rem.String())

		// точное дробное значение в виде рационального числа и вывод с десятичной точностью
		r := new(big.Rat).SetFrac(a, b) // a/b как рациональное число
		// FloatString(n) выводит десятичную запись с n знаками после запятой
		fmt.Println("a / b (точно, рационал) =", r.String())
		fmt.Println("a / b (десятичное, 30 знаков после запятой) =", r.FloatString(30))
	}
}
