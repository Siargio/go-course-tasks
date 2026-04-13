// Задача 4: Композиция интерфейсов
//
// Ожидаемый вывод:
//   engine started
//   engine stopped

package main

import "fmt"

// TODO: объяви интерфейс Starter с методом Start()

// TODO: объяви интерфейс Stopper с методом Stop()

// TODO: объяви интерфейс Machine как композицию Starter и Stopper

// TODO: объяви структуру Engine (без полей)
// TODO: реализуй Start() — выводи "engine started"
// TODO: реализуй Stop()  — выводи "engine stopped"

// TODO: напиши функцию runCycle(m Machine)
//       вызывает m.Start() и m.Stop()

type Starter interface {
	Start()
}

type Stopper interface {
	Stop()
}

type Machine interface {
	Starter
	Stopper
}

type Engine struct{}

func (e Engine) Start() {
	fmt.Println("engine started")
}

func (e Engine) Stop() {
	fmt.Println("engine stopped")
}

func runCycle(m Machine) {
	m.Start()
	m.Stop()
}

func main() {
	// TODO: вызови runCycle с Engine{}

	runCycle(Engine{})
}
