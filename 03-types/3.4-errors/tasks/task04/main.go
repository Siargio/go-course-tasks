// Задача 4: Оборачивание ошибки
//
// Ожидаемый вывод:
//   load data: read file: file not found

package main

import (
	"errors"
	"fmt"
)

// TODO: напиши функцию readFileMock() error
//       возвращай errors.New("file not found")

// TODO: напиши функцию loadData() error
//       вызывает readFileMock() и оборачивает ошибку через fmt.Errorf с %w:
//       "read file: <err>"
//       сама тоже оборачивает: "load data: <err>"

func readFileMock() error {
	return errors.New("file not found")
}

func loadData() error {
	err := readFileMock()
	if err != nil {
		wrapErr := fmt.Errorf("read file: %w", err)
		return fmt.Errorf("load data: %w", wrapErr)
	}
	return nil
}

func main() {
	err := loadData()
	fmt.Println(err)
}
