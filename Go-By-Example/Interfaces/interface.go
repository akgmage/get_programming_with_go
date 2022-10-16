package main

import (
	"fmt"
	"math"
)

/// basic interface for geometric shape
type geometry interface {
	area() float64
	perim() float64
}
// implement geometry interface on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 2, height: 3}
	c := circle{radius: 3}
	measure(r)
	measure(c)
}