package examples

import "fmt"

// SLICES (sequences)
func SlicesExamples() {
	var slice []int					// nil, len == 0
	slice = []int{2, 3, 5}			// [2 3 5], len == 3
	slice = make([]int, 3)			// [0 0 0], len == 3
	slice[1] = 65					// [0 65 0], len == 3
	slice = append(slice, 89)		// [0 65 0 89], len == 4
	slice = slice[1:3]				// [65 0], len == 2
	fmt.Println("slice:", slice)
}
