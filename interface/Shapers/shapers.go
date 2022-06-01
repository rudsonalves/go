package main

import (
	"fmt"
	"math"
)

const PI = math.Pi

// Shapes structs
type Square struct {
	Edge float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Leght  float64
	Height float64
}

// Shapes Areas functions
func (s Square) Area() float64 {
	return math.Pow(s.Edge, 2)
}

func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Leght
}

// Shapes Perimeters functions
func (s Square) Perimeter() float64 {
	return 4 * s.Edge
}

func (c Circle) Perimeter() float64 {
	return 2 * PI * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Leght)
}

// Shapes Strings functinos
func (s Square) String() string {
	return fmt.Sprintf("Square {Edge: %.2f} ", s.Edge)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f} ", c.Radius)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle {Height: %.2f, Legth: %.2f} ", r.Height, r.Leght)
}

// Shapes interfaces
/* Stringer interface is defined in the fmt package like:
type Stringer interface {
	String() string
}
*/

type Shaper interface {
	Area() float64
	Perimeter() float64
}

func TotalArea(shapes ...Shaper) float64 {
	a := 0.0
	for _, s := range shapes {
		a += s.Area()
	}
	return a
}

func TotalPerimeter(shapes ...Shaper) float64 {
	p := 0.0
	for _, s := range shapes {
		p += s.Perimeter()
	}
	return p
}

func main() {
	a := Square{2.5}
	b := Rectangle{1.5, 2.6}
	c := Circle{3.2}

	fmt.Printf("a: %v\nb: %v\nc: %v\n", a, b, c)
	fmt.Printf("Total Area: %.2f\n", TotalArea(a, b, c))
	fmt.Printf("Total Perimeter: %.2f\n", TotalPerimeter(a, b, c))
}
