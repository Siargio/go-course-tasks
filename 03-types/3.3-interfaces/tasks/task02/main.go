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

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	// TODO: создай срез []Shape с Rectangle{Width: 10, Height: 5} и Circle{Radius: 3}
	// TODO: в цикле выведи площадь каждой фигуры через fmt.Printf
	//       для Rectangle: "Rectangle area: %.2f\n"
	//       для Circle:    "Circle area: %.2f\n"
	//       (подсказка: используй type switch или type assertion)

	s := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
	}

	for _, shape := range s {
		switch s := shape.(type) {
		case Rectangle:
			fmt.Printf("Rectangle area: %.2f\n", s.Area())
		case Circle:
			fmt.Printf("Circle area:: %.2f\n", s.Area())
		}
	}

	for _, sh := range s {
		if rect, ok := sh.(Rectangle); ok {
			fmt.Printf("Rectangle area: %.2f\n", rect.Area())
		} else if circle, ok := sh.(Circle); ok {
			fmt.Printf("Circle area:: %.2f\n", circle.Area())
		}
	}
}
