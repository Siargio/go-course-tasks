// Задача 3: Сумма чисел
//
// Ожидаемый вывод:
//   sum int: 15
//   sum float: 7.5

package main

import "fmt"

// TODO: объяви ограничение Number для ~int | ~float64

// TODO: напиши функцию Sum[T Number](items []T) T
//       возвращает сумму всех элементов среза

type Number interface {
	~int | ~float64
}

func Sum[T Number](items []T) T {
	var result T
	for _, item := range items {
		result += item
	}
	return result
}

func main() {
	// TODO: вызови Sum для []int{1, 2, 3, 4, 5} и выведи "sum int: <результат>"
	// TODO: вызови Sum для []float64{1.5, 2.0, 4.0} и выведи "sum float: <результат>"

	sumInt := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("sum int:", sumInt)

	sumFloat := Sum([]float64{1.5, 2.0, 4.0})
	fmt.Println("sum int:", sumFloat)
}
