// Задание 2: Бесконечный генератор чисел Фибоначчи
//
// Числа Фибоначчи: каждое следующее = сумма двух предыдущих.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// Напиши функцию Fibonacci() iter.Seq[int], которая генерирует
// числа Фибоначчи бесконечно.
//
// Внутри итератора:
//   a, b := 0, 1
//   для каждой итерации: yield(a), потом a, b = b, a+b
//
// В main() выведи первые 10 чисел и прерви цикл.
//
// Ожидаемый вывод:
//   0 1 1 2 3 5 8 13 21 34
//
// Запусти: go run main.go

package main

import (
	"fmt"
	"iter"
)

// TODO: напиши функцию Fibonacci() iter.Seq[int]
// Подсказка: используй две переменные a и b
// Не забудь проверить возвращаемое значение yield!
// Функция которая возвращает функцию
func Fibonacci() iter.Seq[int] {
	a, b := 0, 1
	return func(yield func(int) bool) {
		// бесконечный цикл
		for {
			// как только тут будет фолс из-за брейка мы прервем выполнение
			if !yield(a) {
				return
			}
			// вычисляем следующее число
			a, b = b, a+b
		}
	}
}

func main() {
	count := 0
	for n := range Fibonacci() {
		fmt.Print(n, " ")
		count++
		if count == 10 {
			break // прерываем бесконечный итератор
			// как мы в тут сделаем брейк/ретерн/панику или выйдем из области видимости goto
			// это служит флагом фолс для функции yield
		}
	}
	fmt.Println()

	// Пайплайн: 1,2,3,4,5 → квадраты → только чётные
	c := filterEven(square(generate(1, 2, 3, 4, 5)))
	for v := range c {
		fmt.Println(v)
	}
}

// ============================================================
// Задача: Канальный пайплайн  🟢 Junior
// ============================================================
//
// Реализуй трёхступенчатый пайплайн:
//
//   generate(nums...) → square() → filterEven() → вывод
//
//   generate    — принимает срез чисел, отправляет в канал по одному
//   square      — возводит каждое число в квадрат
//   filterEven  — пропускает только чётные числа
//
// Каждая стадия:
//   - принимает <-chan int
//   - возвращает <-chan int
//   - запускает горутину которая закрывает выходной канал когда входной закрыт
//
// Ожидаемый вывод для generate(1,2,3,4,5):
//   4    (2²)
//   16   (4²)
//
// Запуск:
//   go run main.go
//   go test -v ./...

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for val := range in {
			if val%2 == 0 {
				result <- val
			}
		}
	}()
	return result
}
