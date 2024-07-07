package examples

import "fmt"

// SWITCH examples
func SwitchExamples() {
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("I dunno :/")
	}

	switch {
	case i < 0:
		fmt.Println("negative")
	}

	//switch v := i.(type) {
	//case int:
	//	fmt.Println("i is", v)
	//}
}
