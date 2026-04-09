// Задание 2: Счётчик с начальным значением
//
// Напиши функцию makeCounter(start int) func() int, которая:
//   - принимает начальное значение
//   - возвращает функцию-счётчик
//   - каждый вызов возвращает следующее число, начиная с start
//
// Ожидаемый вывод:
//   Счётчик от 5:
//   5
//   6
//   7
//   Счётчик от 100:
//   100
//   101
//   Счётчики независимы - счётчик от 5 продолжает:
//   8
//
// Запусти: go run main.go

package main

import "fmt"

// TODO: напиши функцию makeCounter(start int) func() int
// Подсказка: текущее значение храни в переменной внутри makeCounter,
// и обращайся к ней из возвращаемой функции (это и есть замыкание)

func main() {
	// TODO: создай два независимых счётчика и проверь их работу

	counterFits := makeCounter(5)
	fmt.Println(counterFits())
	fmt.Println(counterFits())
	fmt.Println(counterFits())

	counterSecond := makeCounter(100)
	fmt.Println(counterSecond())
	fmt.Println(counterSecond())
	fmt.Println(counterFits())
}

func makeCounter(start int) func() int {
	// текущее значение счетчика
	current := start

	return func() int {
		// запоминаю текущее значение
		value := current
		// увеличиваю для следующего вызова
		current++
		// вызываю предыдущее значение
		return value
	}
}
