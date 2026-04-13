// Задача 6: Обработка ошибок в цикле
//
// Ожидаемый вывод:
//   ok: 1
//   skip
//   ok: 2
//   skip
//   ok: 3

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO: напиши функцию parseID(s string) (int, error)
//       если s == "", верни 0 и errors.New("empty id")
//       иначе используй strconv.Atoi(s) и верни результат

func parseID(s string) (int, error) {
	if s == "" {
		return 0, errors.New("empty id")
	} else {
		return strconv.Atoi(s)
	}
}

func main() {
	ids := []string{"1", "", "2", "", "3"}

	for _, id := range ids {
		// TODO: вызови parseID(id)
		//       если ошибка — выведи "skip" и продолжи
		//       если успех — выведи "ok: <значение>"
		result, err := parseID(id)

		if err != nil {
			fmt.Println("skip")
			continue
		}

		fmt.Println("ok:", result)

		// _ = id
		// _ = strconv.Atoi // подсказка: используй в parseID
		// fmt.Println("TODO: implement me")
	}
}
