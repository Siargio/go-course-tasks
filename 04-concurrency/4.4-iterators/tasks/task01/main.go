// Задание 1: Итератор для стека
//
// Стек - это структура данных где последний добавленный элемент
// достаётся первым (как стопка тарелок).
//
// Реализуй:
//   type Stack struct { items []int }
//   func (s *Stack) Push(val int)      - добавить элемент наверх
//   func (s *Stack) All() iter.Seq[int] - итератор от вершины к основанию
//
// В main():
//   1. Создай стек и добавь числа 1, 2, 3, 4, 5
//   2. Выведи все элементы через for range
//   3. Выведи элементы пока не встретишь 3 (используй break)
//
// Ожидаемый вывод:
//   Все элементы стека:
//   5
//   4
//   3
//   2
//   1
//
//   До элемента 3:
//   5
//   4
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"iter"
)

// TODO: напиши структуру Stack
type Stack struct {
	items []int
}

// TODO: напиши метод Push(val int)
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// TODO: напиши метод All() iter.Seq[int]
func (s *Stack) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(s.items[i]) {
				return
			}
		}
	}
}

// Подсказка: перебирай items с конца (от len-1 до 0)
// Не забудь проверять возвращаемое значение yield и прерваться если false

func main() {
	// TODO: создай стек, добавь элементы 1..5
	stack := &Stack{}
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	fmt.Println("Все элементы стека:")
	// TODO: for range stack.All() { ... }
	for val := range stack.All() {
		fmt.Println(val)
	}

	fmt.Println("\nДо элемента 3:")
	// TODO: for range stack.All() { if v == 3 { break } ... }
	//stack.All()
	for val := range stack.All() {
		if val == 3 {
			break
		}
		fmt.Println(val)
	}
}
