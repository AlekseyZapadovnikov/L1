/* решим задачу - посчитать самое часто встречаемое слово в трех файлах
используя две реализации конкурентной мапы:
1) RWMutexMap - свою реализацию на базе sync.RWMutex
2) sync.Map - стандартную из библиотеки sync

замеры времени и сравнение результатов
*/



package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"
)

// --------------------------------------------------------------------
// 1. реализация: RWMutexMap
// --------------------------------------------------------------------

type RWMutexMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewRWMutexMap() *RWMutexMap {
	return &RWMutexMap{m: make(map[string]int)}
}

func (r *RWMutexMap) Inc(key string) {
	r.mu.Lock()
	r.m[key]++
	r.mu.Unlock()
}

func (r *RWMutexMap) Range(f func(key string, value int)) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for k, v := range r.m {
		f(k, v)
	}
}

// --------------------------------------------------------------------
// 2. Воркеры для каждой мапы
// --------------------------------------------------------------------

// countWordsInFileRW - воркер для RWMutexMap.
func countWordsInFileRW(path string, counts *RWMutexMap, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка открытия файла %s: %v", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		cleanWord := cleanString(scanner.Text())
		if len(cleanWord) > 0 {
			counts.Inc(cleanWord)
		}
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Printf("Ошибка чтения файла %s: %v", path, err)
	}
}

// countWordsInFileSM - воркер для sync.Map.
func countWordsInFileSM(path string, counts *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Ошибка открытия файла %s: %v", path, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		cleanWord := cleanString(scanner.Text())
		if len(cleanWord) > 0 {
			ptr, _ := counts.LoadOrStore(cleanWord, new(int64))
			atomic.AddInt64(ptr.(*int64), 1)
		}
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Printf("Ошибка чтения файла %s: %v", path, err)
	}
}

// cleanString - вспомогательная функция для очистки слов.
func cleanString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) || unicode.IsSpace(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
}

// --------------------------------------------------------------------
// 3. Главная функция с бенчмарком
// --------------------------------------------------------------------

func main() {
	paths := []string{"./L1.7/static/file1.txt", "./L1.7/static/file2.txt", "./L1.7/static/file3.txt"}

	// --- Тест 1: RWMutexMap ---
	fmt.Println("\n--- Запуск с RWMutexMap ---")
	rwMutexMap := NewRWMutexMap()
	var wg1 sync.WaitGroup
	start1 := time.Now()
	for _, path := range paths {
		wg1.Add(1)
		go countWordsInFileRW(path, rwMutexMap, &wg1)
	}
	wg1.Wait()
	maxCount1 := 0
	rwMutexMap.Range(func(key string, value int) {
		if value > maxCount1 {
			maxCount1 = value
		}
	})
	duration1 := time.Since(start1)
	fmt.Printf("Самое частое слово встречается %d раз(а)\n", maxCount1)
	fmt.Printf("Время выполнения: %v\n", duration1)

	// --- Тест 2: Стандартная sync.Map ---
	fmt.Println("\n--- Запуск со стандартной sync.Map ---")
	var syncMap sync.Map
	var wg2 sync.WaitGroup
	start2 := time.Now()
	for _, path := range paths {
		wg2.Add(1)
		go countWordsInFileSM(path, &syncMap, &wg2)
	}
	wg2.Wait()
	maxCount2 := 0
	syncMap.Range(func(key, value interface{}) bool {
		count := int(atomic.LoadInt64(value.(*int64)))
		if count > maxCount2 {
			maxCount2 = count
		}
		return true
	})
	duration2 := time.Since(start2)
	fmt.Printf("Самое частое слово встречается %d раз(а)\n", maxCount2)
	fmt.Printf("Время выполнения: %v\n", duration2)
}

/*
PS C:\Users\Asus\Desktop\go_p\WB_L1!> go run .\L1.7\main.go

--- Запуск с RWMutexMap ---
Самое частое слово встречается 18 раз(а)
Время выполнения: 1.0952ms

--- Запуск со стандартной sync.Map ---
Самое частое слово встречается 18 раз(а)
Время выполнения: 523.8µs
PS C:\Users\Asus\Desktop\go_p\WB_L1!> 

------------------------------------------------------------------------
Выводы: я написал фигню, потому что по-хорошему блокировать только нужное значение, а не всю мапу
поэтому sync.Map оказался быстрее, я думаю
надо было не int использовать, а *int64 и атомарно инкрементить, но я это понял уже после того как всё написал
------------------------------------------------------------------------
*/