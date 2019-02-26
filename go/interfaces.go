package main

import "fmt"
import "math"

func main() {

	rect := Rectangle{20,50}
	circ := Circle{4}

	fmt.Println("Rectangle area = ", getArea(rect))
	fmt.Println("Circle area = ", getArea(circ))

}

// an interface defines a list of methods a type MUST implement
// if that type implements those methods, the proper method is executed
// -even if the original is referred to with the interface name
type Shape interface {
	area() float64
}

// make normal struct objects
type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}


// create methods that accept those structs for each type
func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2) 
}

// method that works on any "shape"
// can accept any "shape" object
func getArea(shape Shape) float64{
	return shape.area()
}
