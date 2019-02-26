package main

import "fmt"

func main() {

	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 10}

	fmt.Println("Rectangle is", rect1.width, "wide")

	fmt.Println("Rectangle area is ", rect1.area())

}

// lets make an object type of our own
// type TypeName struct { field type; field type; field type; }
type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

// attach a function to our custom type, rectangle
// func (arg *TypeToAttachTo) funcName() returnType {}
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
