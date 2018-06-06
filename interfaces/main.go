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

func (s *rectangle) GetArea() float64 {
	return s.length * s.width
}

func (s *rectangle) Name() string {
	return "Rectangle"
}

func (s *rectangle) ToDisplayString() string {
	return fmt.Sprintf("%v * %v", s.length, s.width)
}

func main() {
	shapes := []ShapeAreaCalculator{}

	s1 := &triangle{bottom: 10, height: 8}
	s2 := &circle{radius: 8}
	s3 := &rectangle{length: 12, width: 5}

	shapes = append(shapes, s1, s2, s3)

	for _, shape := range shapes {
		fmt.Printf("  %s area (%s) = %v\n",
			shape.Name(),
			shape.ToDisplayString(),
			shape.GetArea())
	}
}
