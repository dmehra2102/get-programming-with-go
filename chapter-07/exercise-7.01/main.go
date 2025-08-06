package main

import (
	"fmt"
	"math"
)

type Gemetry interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Measure(g Gemetry){
	fmt.Printf("Shape Type: %T\n", g)
	fmt.Printf("Area: %.2f\n", g.Area())
	fmt.Printf("Perimeter: %.2f\n", g.Perimeter())
	fmt.Println("---")
}

func main(){
	rect := Rectangle{Width: 3, Height: 4}
	circle := Circle{Radius: 5}

	fmt.Println("Measuring Rectangle:")
	Measure(rect)

	fmt.Println("Measuring Circle:")
	Measure(circle)
}