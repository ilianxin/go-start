package main

import "fmt"

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	a, b float32
}
type Circle struct {
	a float32
}

func (r Rectangle) Area() float32 {
	return r.a * r.b
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.a + r.b)
}

func (c Circle) Area() float32 {
	return 3.14 * c.a * c.a
}

func (c Circle) Perimeter() float32 {
	return 2 * 3.14 * c.a
}

func main() {
	rec := Rectangle{3, 4}
	cir := Circle{3}
	fmt.Printf("Rectangle Area: %f, Perimeter: %f\n", rec.Area(), rec.Perimeter())
	fmt.Printf("Circle Area: %v, Perimeter: %v", cir.Area(), cir.Perimeter())
}
