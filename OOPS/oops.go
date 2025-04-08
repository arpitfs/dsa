package oops

import (
	"fmt"
	"math"
)

// Abstraction: Define a Shape interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Encapsulation + Composition: Define a Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Polymorphism: a function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println("---")
}
func OOPS() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	fmt.Println("Circle:")
	printShapeInfo(circle)

	fmt.Println("Rectangle:")
	printShapeInfo(rectangle)
}
