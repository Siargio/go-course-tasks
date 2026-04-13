// Задача 4: Поиск индекса
//
// Ожидаемый вывод:
//   index: 2
//   index: -1

package main

import "fmt"

// TODO: напиши функцию IndexOf[T comparable](items []T, target T) int
//       возвращает индекс target в items, или -1 если не найден

func IndexOf[T comparable](items []T, target T) int {
	var result int = -1

	for i, v := range items {
		if v == target {
			result = i
			return result
		}
	}

	return -1
}

func main() {
	// TODO: вызови IndexOf([]string{"a", "b", "c"}, "c") и выведи "index: <результат>"
	// TODO: вызови IndexOf([]string{"a", "b", "c"}, "z") и выведи "index: <результат>"
	fmt.Println("TODO: implement me")

	fmt.Println("index:", IndexOf([]string{"a", "b", "c"}, "c"))
	fmt.Println("index:", IndexOf([]string{"a", "b", "c"}, "z"))
}
