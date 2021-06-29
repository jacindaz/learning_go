package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}
type Square struct {
	Width float64
	Height float64
}
type Sizer interface {
	Area() float64
}
type Shaper interface {
	Sizer
	fmt.Stringer
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (s Square) Area() float64 {
	return s.Width * s.Height
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {Radius: %.2f}", c.Radius)
}

func (s Square) String() string {
	return fmt.Sprintf("Square {Width: %.2f, Height: %.2f}", s.Width, s.Height)
}

func PrintArea(s Shaper) {
	fmt.Printf("area of %s is %.2f\n", s.String(), s.Area())
}

func main() {
	circle := Circle{Radius: 10}
	square := Square{Height: 10, Width: 5}
	PrintArea(circle)
	PrintArea(square)

	less := Less(circle, square)
	fmt.Printf("%+v is the smallest\n", less)
}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s2
	}

	return s2
}
