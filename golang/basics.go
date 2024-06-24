package main

import "fmt"

func main() {
	values_examples()
	variables_examples()
	constants_examples()
	loops_examples()
	conditions_examples()
	switch_examples()
	arrays_examples()
	slices_examples()
	maps_examples()
	range_examples()
	functions_examples()
	pointers_examples()
}

/** Print VALUES examples */
func values_examples() {
	fmt.Println("Hello", "World!")	// Hello World!
	fmt.Println("1 + 1 =", 1 + 1)	// 1 + 1 = 2
	fmt.Println("7.0 / 3.0 =", 7.0 / 3.0)
	fmt.Println(true && false)		// false
	fmt.Println(!false)				// true
}

/** VARIABLES examples */
func variables_examples() {
	var a = 0		// var a int = 0
	var b int = 0	// var b int = 0
	c := 0			// var c int = 0

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

/** CONSTANTS examples */
func constants_examples() {
	const d = 3e10	// const d int = 30000000000
	fmt.Println("d:", d)
}

/** LOOPS exmaples */
func loops_examples() {
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

/** CONDITIONS examples */
func conditions_examples() {
	if n := 1 + 1; n == 2 {
		fmt.Println("Right")
	} else if n < 2 || n > 2 {
		fmt.Println("Wrong")
	}
}

/** SWITCH examples */
func switch_examples() {
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

/** ARRAYS examples */
func arrays_examples() {
	integers := [...]int{3: 30, 40, 60}	// [5]int[0, 0, 30, 40, 60}
	fmt.Println(integers)
}

/** SLICES (sequences) */
func slices_examples() {
	var slice []int					// nil, len == 0
	slice = []int{2, 3, 5}			// [2 3 5], len == 3
	slice = make([]int, 3)			// [0 0 0], len == 3
	slice[1] = 65					// [0 65 0], len == 3
	slice = append(slice, 89)		// [0 65 0 89], len == 4
	slice = slice[1:3]				// [65 0], len == 2
	fmt.Println("slice:", slice)
}

/** MAPs examples */
func maps_examples() {
	var mapp map[string]int			// nil
	mapp = make(map[string]int)		// map[]
	mapp = map[string]int{"two":2}	// map[two: 2]
	mapp["one"] = 1					// map[one:1 two:2]
	delete(mapp, "two")				// map[one:1]
	clear(mapp)						// map[]
	fmt.Println("mapp:", mapp)
}

/** Range iterates over elements in a variety of data structures. */
func range_examples() {
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

/** FUNCTIONS examples */
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
func functions_examples() {
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

/** POINTERS examples */
func pointers_examples() {
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
