package main

import (
	"fmt"
	"math"
)


func main(){
	var circle Circle
	circle.diameter = 8

	fmt.Println(circle.diameter)

	triangle := new(Triangle)
	triangle.x = 8

	circle1 := Circle{diameter: 7, radius: 4}

	fmt.Printf("circle1.diameter: %v\n", circle1.diameter)


	fmt.Printf("circleArea(circle): %v\n", circleArea(circle1))


	circle.area()

	sqr := Square{side: 5}

	fmt.Printf("square area: %v\n", sqr.area())

}


type Circle struct{
	diameter int8
	radius int8

}

type Triangle struct{
	x, y, z float64
}


func circleArea(c Circle) float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

//
func (c Circle) area() float64 {
	return math.Acos(float64(c.diameter))
}


// is-a relationship
type Polygon struct{
	Circle
	Triangle
}

//  has-a relationship
type Omo struct {
	circle Circle
	triangle Triangle
}


// interface
type Shape interface {
	area() float64
}

type Square struct {
	Shape
	side int8
}

func (s Square) area() float64 {
	return math.Pow(float64(s.side), 2)
}
