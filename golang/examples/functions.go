package examples

import "fmt"

func sum(a int, b int) int {
	return a + b
}
func vals() (int, string) {
	return 5, "value"
}
func total(n ...int) int {
	t := 0
	for _, v := range n {
		t += v
	}
	return t
}

// FUNCTIONS examples
func FunctionsExamples() {
	fmt.Println("func:", sum(5, 6)) // 11
	i, v := vals()
	fmt.Printf("func: %d %s\n", i, v) // 5 value

	fmt.Println("func:", total(8, 5, 7)) // 20
	s := []int{3, 4, 5}
	fmt.Println("func:", total(s...)) // 12

	sum := func(a, b int) int {
		fmt.Println("OK!")
		return a + b
	}
	fmt.Println(sum(3, 4)) // sum is the anonymous func, not the global.
}
