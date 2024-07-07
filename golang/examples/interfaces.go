package examples

import "fmt"

type geometry interface {
	area() float64
}
type rect struct {
	width, height float64
}
func (r rect) area() float64 {
	return r.width * r.height
}

// Interfaces: named collections of method signatures.
func InterfacesExamples() {
	measure := func(g geometry) {
		fmt.Println("geometry:", g)
		fmt.Println("geometry area:", g.area())
	}
	measure(rect{4.5, 10})
}
