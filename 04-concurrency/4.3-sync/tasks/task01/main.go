// Задание 1: Безопасный счётчик через Mutex
//
// Создай структуру SafeCounter с:
//   - полем mu типа sync.Mutex
//   - полем value типа int
//   - методом Increment() - увеличивает value на 1 (с блокировкой)
//   - методом Value() int - возвращает value (с блокировкой)
//
// Запусти 1000 горутин, каждая вызывает counter.Increment().
// После wg.Wait() выведи финальное значение.
//
// Проверь отсутствие гонок:
//   go run -race main.go
//
// Ожидаемый вывод:
//   Финальный счётчик: 1000
//   (и никаких WARNING: DATA RACE)
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"sync"
)

// TODO: напиши структуру SafeCounter
// type SafeCounter struct { ... }
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// TODO: напиши метод Increment()
func (s *SafeCounter) Increment() {
	//блочим
	s.mu.Lock()
	// освобождаем
	defer s.mu.Unlock()
	//только с одного потока пишем
	s.value++
}

// TODO: напиши метод Value() int
func (s *SafeCounter) Value() int {
	//блочим
	s.mu.Lock()
	//освобождаем
	defer s.mu.Unlock()
	//только с одного потока читам
	return s.value
}

func main() {
	// TODO: создай SafeCounter и запусти 1000 горутин
	// Каждая горутина вызывает counter.Increment()
	// После завершения всех горутин выведи counter.Value()

	// создал структуру
	s := &SafeCounter{}
	// создал группу для отслеживания
	wg := &sync.WaitGroup{}

	// счетчик до тысячи
	for i := 0; i < 1000; i++ {
		//счетчик, одна горутина
		wg.Add(1)
		//создаем горутинку
		go func() {
			// не забываем уменьшит счетчик
			defer wg.Done()
			//вызываем наш метод +1
			s.Increment()
		}()
	}
	// ждем пока счетчик не станет 0
	wg.Wait()
	fmt.Println("Финальный счётчик:", s.Value())
}
