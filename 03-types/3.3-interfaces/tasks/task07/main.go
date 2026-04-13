// Задача 7: Интеграционная задача — система логирования
//
// Ожидаемый вывод:
//   processing order order-1
//   validating order
//   order done
//   [INFO] processing order order-1
//   [INFO] validating order
//   [INFO] order done

package main

import "fmt"

// TODO: объяви интерфейс Logger с методом Log(message string)

// TODO: объяви структуру ConsoleLogger (без полей)
// TODO: реализуй Log(message string) — просто выводи message

// TODO: объяви структуру PrefixLogger с полем Prefix (string)
// TODO: реализуй Log(message string) — выводи "<Prefix> <message>"

// TODO: напиши функцию processOrder(logger Logger, id string)
//       пишет 3 сообщения:
//       "processing order <id>"
//       "validating order"
//       "order done"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type PrefixLogger struct {
	Prefix string
}

func (p PrefixLogger) Log(message string) {
	fmt.Println(p.Prefix, message)
}

func processOrder(logger Logger, id string) {
	logger.Log("processing order " + id)
	logger.Log("validating order")
	logger.Log("order done")
}

func main() {
	// TODO: вызови processOrder с ConsoleLogger{} и id "order-1"
	// TODO: вызови processOrder с PrefixLogger{Prefix: "[INFO]"} и id "order-1"

	processOrder(ConsoleLogger{}, "order-1")
	processOrder(PrefixLogger{Prefix: "[INFO]"}, "order-1")
}
