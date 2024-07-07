package examples

import "fmt"

// Range iterates over elements in a variety of data structures.
func RangeExamples() {
	slice := []int{5, 8, 1, 10}
	var sum int // int 0
	for i, v := range slice {
		if i == 1 {
			fmt.Println("Ignoring second value which is: ", v)
			continue
		}
		fmt.Printf("%d - Iterating over VALUES of slice.\n", i)
		sum += v
	}
	fmt.Println("sum:", sum)

	for i := range slice {
		fmt.Printf("%d - Iterating over INDEXES of slice.\n", i)
	}
}
