package examples

import "fmt"

type person struct {
	name string
	age uint8
}
func (p person) greetig() {
	fmt.Printf("My name is %s and I have %d years old.\n", p.name, p.age)
}
func (p *person) incrementAge() {
	p.age++
}

// Structs: typed collections of fields.
// Receiver function: method that belongs to a single non-variadic type.
func StructsExamples() {
	p1 := person{name: "João", age: 54}
	p2 := &p1
	p2.age = 21
	fmt.Println("person1:", p1)

	p1.incrementAge()
	p1.greetig()

	p3 := struct {
		name string
		hairColor string
	} {
		"Pedro",
		"Azul",
	}
	fmt.Println("person3:", p3)
}
