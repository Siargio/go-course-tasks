// Задача: Срез интерфейсов
//
// Ожидаемый вывод:
//   Rectangle area: 50.00
//   Circle area: 28.27

package main

import (
	"fmt"
	"math"
)

// TODO: объяви интерфейс Shape с методом Area() float64

// TODO: объяви структуру Rectangle с полями Width, Height float64
// TODO: реализуй метод Area() float64 для Rectangle

// TODO: объяви структуру Circle с полем Radius float64
// TODO: реализуй метод Area() float64 для Circle (используй math.Pi)

func main() {
	// TODO: создай срез []Shape с Rectangle{Width: 10, Height: 5} и Circle{Radius: 3}
	// TODO: в цикле выведи площадь каждой фигуры через fmt.Printf
	//       для Rectangle: "Rectangle area: %.2f\n"
	//       для Circle:    "Circle area: %.2f\n"
	//       (подсказка: используй type switch или type assertion)
	_ = math.Pi
	fmt.Println("TODO: implement me")
}
