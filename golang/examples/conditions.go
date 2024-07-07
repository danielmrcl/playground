package examples

import "fmt"

// CONDITIONS examples
func ConditionsExamples() {
	if n := 1 + 1; n == 2 {
		fmt.Println("Right")
	} else if n < 2 || n > 2 {
		fmt.Println("Wrong")
	}
}
