package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func worker(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// читаем из канала
	for val := range ch {
		fmt.Printf("worker %d прочитал %d\n", id, val)
	}
	fmt.Printf("worker %d: завершаюсь\n", id)
}

func main() {
	workerAmount, err := strconv.Atoi(os.Args[len(os.Args)-1]) // парсим последний аргумент командной строки
	// проводим валидацию числа
	if err != nil || workerAmount <= 0 {
		fmt.Println("Вы ввели невалидное число")
		return
	}

	// Ловим Ctrl+C
	sigCh := make(chan os.Signal, 1) // канал сигналов
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	/* ignal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM) говорит рантайму Go — 
	перенаправлять поступающие ОС-сигналы SIGINT и SIGTERM в канал sigCh*/

	ch := make(chan int)
	wg := sync.WaitGroup{}

	// Запускаем воркеров
	wg.Add(workerAmount)
	for i := 1; i <= workerAmount; i++ {
		go worker(ch, i, &wg)
	}

	// Главная горутина: продюсер + ожидание сигнала
	counter := 0

// синтаксическиц сахар, это то чего мне так не хватала в питоне (двойной брэйк) <3
loop:
	for {
		select {
		case <-sigCh: // если пришёл Ctrl+C
			fmt.Println("\nполучен Ctrl+C, закрываю канал...")
			close(ch)
			break loop
		default:
			counter++
			ch <- counter
		}
	}

	// Дожидаемся завершения воркеров
	wg.Wait()
	fmt.Println("все воркеры завершились, завершаем программу...")
}

/* 
я выбрал способ завершения горутин с помощью закрытия канала, тк

Во-первых, это позволяет реализовать логику корректного закрытия горутин
я имею в виду, что мы не просто закрываем все горутины,
а дочитываем все данные из канала и только потом закрываем

Во-вторых, я не понял как реализовать подобную логику завершения используя context
он ведь hard завершает все горутины и всё, не понял как мне дочитать оставшиеся данные

в моей программе горутины дочитывают данные из канала, потому что 
1) есть WaitGroup
2) for val := range ch будет работать пока в канале есть данные
*/