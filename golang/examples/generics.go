package examples

import "fmt"

type Calculable interface {
	int | float32 | float64
}

func DoubleMe[T Calculable](v T) T {
	return v * 2
}

// Provide generic type examples implementation.
func GenericsExamples() {
	fmt.Printf("(generics) result: %.2f\n", DoubleMe(5.2))
	fmt.Printf("(generics) result: %d\n", DoubleMe(25))
	//fmt.Printf("(generics) result: %d\n", DoubleMe(uint(25)))
	fmt.Println("(generics) Done!")
}
