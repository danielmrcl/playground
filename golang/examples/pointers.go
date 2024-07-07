package examples

import "fmt"

// POINTERS examples
func PointersExamples() {
	zeroval := func(ival int) { // passed by value
		ival = 0 // change just ival
	}
	zeroptr := func(iptr *int) { // passed by memory location
		*iptr = 0 // change value in memory
	}

	i := 1
	fmt.Println("initial:", i) // 1
	zeroval(i)
	fmt.Println("zeroval:", i) // 1
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0
    fmt.Println("pointer:", &i) // 0xc000112038
}
