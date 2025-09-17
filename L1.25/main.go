package main

import (
	"context"
	"fmt"
	"time"
)

/* имитирую таймер через контекст с таймаутом,
под копотом там где-то используется таймер
вот такой (это из пакета context)
type timerCtx struct {
	cancelCtx
	timer *time.Timer // Under cancelCtx.mu.

	deadline time.Time
}
но на всякий я написал свою реализацию с таймером */

func sleep(t time.Duration) {
	// Отнимаем погрешность таймера
	t -= 18 * time.Millisecond

	// Если итоговая длительность стала отрицательной, ничего не делаем
	if t < 0 {
		t = 0
	}

	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	<-ctx.Done()
}

// реализация с таймером
func sleep2(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	startTime := time.Now()
	fmt.Println("Start:", startTime.Format("15:04:05.000"))

	sleep(3 * time.Second)

	endTime := time.Now()
	fmt.Println("End:  ", endTime.Format("15:04:05.000"))
	
	fmt.Printf("Total duration: %v\n", endTime.Sub(startTime))
}