package main

import (
	"fmt"
	"reflect"
)

func detectType(x interface{}) string {
	if x == nil {
		return "nil"
	}

	// сначала — type switch для конкретных типов
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("int (%d)", v)
	case string:
		return fmt.Sprintf("string (%q)", v)
	case bool:
		return fmt.Sprintf("bool (%t)", v)
	}

	// если не один из выше — используем reflect
	t := reflect.TypeOf(x)
	if t.Kind() == reflect.Chan {
		// Узнаём направление канала и тип его элементов
		var dir string
		switch t.ChanDir() {
		case reflect.SendDir:
			dir = "send-only"
		case reflect.RecvDir:
			dir = "recv-only"
		case reflect.BothDir:
			dir = "bidirectional"
		}
		return fmt.Sprintf("chan (%s of %s)", dir, t.Elem().String())
	}

	return fmt.Sprintf("other (%s)", t.String())
}

func main() {
	ch := make(chan int)                // bidirectional chan of int
	var sendOnly chan<- int = ch        // send-only view
	var recvOnly <-chan int = ch        // recv-only view
	var nilChan chan int                // nil channel
	var unknown = 3.14                  // float64 — для демонстрации "other"

	values := []interface{}{
		42,
		"hello",
		true,
		ch,
		sendOnly,
		recvOnly,
		nilChan,
		nil,
		unknown,
	}

	for _, v := range values {
		fmt.Println(detectType(v))
	}
}
