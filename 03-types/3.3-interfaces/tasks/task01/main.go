// Задача: Базовый интерфейс
//
// Ожидаемый вывод:
//   Hello, Anna
//   Welcome, stranger

package main

import "fmt"

// TODO: объяви интерфейс Greeter с методом Greet() string

// TODO: объяви структуру User с полем Name string
// TODO: реализуй метод Greet() string для User — возвращай "Hello, " + u.Name

// TODO: объяви структуру Guest (без полей)
// TODO: реализуй метод Greet() string для Guest — возвращай "Welcome, stranger"

// TODO: напиши функцию printGreeting(g Greeter), которая печатает g.Greet()

type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

type Guest struct{}

func (g Guest) Greet() string {
	return "Welcome, stranger"
}

func printGreeting(g Greeter) string {
	return g.Greet()
}

func main() {
	// TODO: вызови printGreeting для User{Name: "Anna"} и Guest{}

	fmt.Println(printGreeting(User{Name: "Anna"}))
	fmt.Println(printGreeting(Guest{}))
}
