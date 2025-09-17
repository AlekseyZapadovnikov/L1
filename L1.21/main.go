package main

import (
	"fmt"
	"time"
)

// Существующий (legacy) код, который мы не можем/не хотим менять
type LegacyLogger struct{}

func (l *LegacyLogger) Log(msg string) {
	// Простая имитация старого логгера: вывод с временной меткой
	fmt.Printf("%s | %s\n", time.Now().Format(time.RFC3339), msg)
}

// Новый интерфейс, который ожидает клиент
type AppLogger interface {
	Info(msg string)
	Error(msg string)
}

// Адаптер: делает LegacyLogger совместимым с AppLogger
type LegacyLoggerAdapter struct {
	adaptee *LegacyLogger
}

/* Компиляторная проверка: LegacyLoggerAdapter действительно реализует AppLogger
это nil-указатель, явно приведённый к типу *LegacyLoggerAdapter.
То есть выражение имеет тип *LegacyLoggerAdapter
var _ AppLogger = (*LegacyLoggerAdapter)(nil) */

func (a *LegacyLoggerAdapter) Info(msg string) {
	// Можно добавить любую логику преобразования/форматирования
	a.adaptee.Log("[INFO] " + msg)
}

func (a *LegacyLoggerAdapter) Error(msg string) {
	a.adaptee.Log("[ERROR] " + msg)
}

// Клиентский код, который ожидает интерфейс AppLogger
func ProcessSomething(logger AppLogger) {
	logger.Info("start processing")
	// какая-то работа
	logger.Error("something went wrong")
	logger.Info("end processing")
}

func main() {
	legacy := &LegacyLogger{}

	// Создаём адаптер, который делает legacy совместимым с AppLogger
	adapter := &LegacyLoggerAdapter{adaptee: legacy}

	// Передаём адаптер в клиентский код
	ProcessSomething(adapter)
}

/* Когда применять патерн адаптер? 
1) Нужно интегрировать существующий код (legacy, сторонняя библиотека, драйвер)
   с новым API, и менять либо старый код, либо клиента нельзя/невыгодно.
2) Интерфейсы несовместимы: сигнатуры/семантика методов
   отличаются — адаптер преобразует вызовы/форматы.

Плюсы/ Минусы
Плюсы:
1) нет необходимости переписывать уже рабочую реализацию.
2) преобразования сосредоточены в одном месте. (меняем только новую структуру Adapter)
3) Независимость разработки — клиент и adaptee могут эволюционировать автономно.

Минусы: 
1) немного усложняет трассировку (требует читать ещё один класс/файл).
2) если адаптер содержит сложную логику маппинга, он становится тяжёлым

Реальные примеры использования.
1) Есть библиотечный http.HandlerFunc, а мы хотим использовать свой интерфейс AppHandler.
   Пишем адаптер func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)
2) Паттерн «Адаптер» применяется в FastCGI, потому что библиотека переводит
   низкоуровневые FCGI-рекорды и потоки (stdin/stdout, заголовки, параметры окружения) в интерфейс,
   ожидаемый приложением (например `http.Request` и `http.ResponseWriter`),
   позволяя повторно использовать стандартные HTTP-хендлеры без изменения их кода.
*/
