package main

import "fmt"

func main() {
	/* VALUES */
	fmt.Println("Hello", "World!")	// Hello World!
	fmt.Println("1 + 1 =", 1 + 1)	// 1 + 1 = 2
	fmt.Println("7.0 / 3.0 =", 7.0 / 3.0)
	fmt.Println(true && false)		// false
	fmt.Println(!false)				// true

	/* VARIABLES */
	var a = 0		// var a int = 0
	var b int = 0	// var b int = 0
	c := 0			// var c int = 0

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	/* CONSTANTS */
	const d = 3e10	// const d int = 30000000000

	fmt.Println("d:", d)

	/* LOOPS */
	i := 1
	for i <= 3 {
		fmt.Println("i:", i)
		i = i + 1
	}

	for {
		fmt.Println("loop")
		break
	}

	/* CONTITIONS */
	if n := 1 + 1; n == 2 {
		fmt.Println("Right")
	} else if n < 2 || n > 2 {
		fmt.Println("Wrong")
	}

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

	/* ARRAYS */
	integers := [...]int{3: 30, 40, 60}	// [5]int[0, 0, 30, 40, 60}
	fmt.Println(integers)

	/* SLICES */
	var s = make([]int, 3)
	
	// TODO
}
