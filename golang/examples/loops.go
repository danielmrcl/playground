package examples

import "fmt"

// LOOPS exmaples
func LoopsExamples() {
	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i = i + 1
	}

	for {
		fmt.Println("loop")
		break
	}
}
