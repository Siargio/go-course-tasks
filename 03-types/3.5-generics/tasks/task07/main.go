// Задача 7: Интеграционная задача — обобщенный Store
//
// Ожидаемый вывод:
//   strings: [go rust python]
//   ints: [1 2 3]
//   contains "go": true
//   contains "java": false
//   contains 2: true
//   contains 5: false

package main

import "fmt"

// TODO: объяви структуру Store[T any] с полем items []T
// TODO: добавь метод Add(item T) — добавляет элемент в items
// TODO: добавь метод All() []T — возвращает items

// TODO: напиши функцию Contains[T comparable](items []T, target T) bool
//       возвращает true если target есть в items

type Store[T any] struct {
	items []T
}

func (s *Store[T]) Add(item T) {
	s.items = append(s.items, item)
}

func (s Store[T]) All() []T {
	return s.items
}

func Contains[T comparable](items []T, target T) bool {
	for _, v := range items {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	// TODO: создай Store[string], добавь "go", "rust", "python"
	//       выведи "strings:", ss.All()

	// TODO: создай Store[int], добавь 1, 2, 3
	//       выведи "ints:", si.All()

	// TODO: выведи результаты Contains:
	//       Contains(ss.All(), "go")  → "contains \"go\": true"
	//       Contains(ss.All(), "java") → "contains \"java\": false"
	//       Contains(si.All(), 2)     → "contains 2: true"
	//       Contains(si.All(), 5)     → "contains 5: false"

	var storeString Store[string]
	storeString.Add("go")
	storeString.Add("rust")
	storeString.Add("python")
	fmt.Println("strings:", storeString.All())

	var storeInt Store[int]
	storeInt.Add(1)
	storeInt.Add(2)
	storeInt.Add(3)
	fmt.Println("ints:", storeInt.All())

	fmt.Println("contains go:", Contains(storeString.All(), "go"))
	fmt.Println("contains java:", Contains(storeString.All(), "java"))
	fmt.Println("contains 2:", Contains(storeInt.All(), 2))
	fmt.Println("contains 5:", Contains(storeInt.All(), 5))
}
