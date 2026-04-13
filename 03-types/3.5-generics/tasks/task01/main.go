// Задача: Универсальный возврат
//
// Ожидаемый вывод:
//   42
//   hello
//   true

package main

import "fmt"

// TODO: напиши функцию Echo[T any](v T) T
//       она должна возвращать переданное значение без изменений

func Echo[T any](v T) T {

	return v
}

func main() {
	// TODO: вызови Echo с int (42), string ("hello"), bool (true)
	//       и выведи каждый результат через fmt.Println

	echo := Echo([]any{42, "hello", true})

	for _, result := range echo {
		fmt.Println(result)
	}

}
