package main

import (
	"fmt"
	"math"
)

type ShapeAreaCalculator interface {
	GetArea() float64
	Name() string
	ToDisplayString() string
}

// Type triangle definitions
type triangle struct {
	bottom float64
	height float64
}

// Following functions are the implementation of interface ShapeAreaCalculator
// for type triangle
func (s *triangle) GetArea() float64 {
	return s.bottom * s.height / 2
}

func (s *triangle) Name() string {
	return "Triangle"
}

func (s *triangle) ToDisplayString() string {
	return fmt.Sprintf("%v * %v / 2", s.bottom, s.height)
}

type circle struct {
	radius float64
}

// Following functions are the implementation of interface ShapeAreaCalculator
// for type circle
func (s *circle) GetArea() float64 {
	return math.Pi * s.radius * s.radius
}

func (s *circle) Name() string {
	return "Circle"
}

func (s *circle) ToDisplayString() string {
	return fmt.Sprintf("Pi * %v * %v", s.radius, s.radius)
}

type rectangle struct {
	length float64
	width  float64
}

// Following functions are the implementation of interface ShapeAreaCalculator
// for type rectangle
func (s *rectangle) GetArea() float64 {
	return s.length * s.width
}

func (s *rectangle) Name() string {
	return "Rectangle"
}

func (s *rectangle) ToDisplayString() string {
	return fmt.Sprintf("%v * %v", s.length, s.width)
}

type square struct {
	width float64
}

// Following functions are the implementation of interface ShapeAreaCalculator
// for type square
func (s *square) GetArea() float64 {
	return s.width * s.width
}

func (s *square) Name() string {
	return "Square"
}

func (s *square) ToDisplayString() string {
	return fmt.Sprintf("%v * %v", s.width, s.width)
}

func main() {

	shapes := []ShapeAreaCalculator{}

	// Why we can use the address of a specific typed struct that
	// implements ShapeAreaCalculator interface
	s1 := &triangle{bottom: 10, height: 8}
	s2 := &circle{radius: 8}
	s3 := &rectangle{length: 12, width: 5}
	s4 := &square{width: 5}

	shapes = append(shapes, s1, s2, s3, s4)

	sumShapeArea(shapes)
}

// Since we use ShapeAreaCalculator interface as the representation of a shape for the
// method, any struct that implements ShapeAreaCalculator interface can be used in this
// method. When adding new shape, this method will also work without any change.
func sumShapeArea(shapes []ShapeAreaCalculator) {
	var sumOfArea float64
	for _, shape := range shapes {
		area := shape.GetArea()
		sumOfArea += area
		fmt.Printf("  Area of %s \t(%s) = %v\n",
			shape.Name(),
			shape.ToDisplayString(),
			area)
	}

	fmt.Printf("The sum of area of %d shapes are %v\n", len(shapes), sumOfArea)
}
