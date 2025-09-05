// Плюсы: позволяет использовать «несовместимый» код без переписывания.
// Минусы: добавляет слой абстракции и может немного усложнить код.
package main

import "fmt"

type OldLogger struct{}

func (l *OldLogger) PrintLog(message string) {
	fmt.Println("OLD LOG:", message)
}

type Logger interface {
	Info(msg string)
}

type LoggerAdapter struct {
	old *OldLogger
}

func (a *LoggerAdapter) Info(msg string) {
	fmt.Println(msg)
}

func main() {
	old := &OldLogger{}
	adapter := &LoggerAdapter{old: old}

	var logger Logger = adapter
	logger.Info("Hello Adapter!")
}
