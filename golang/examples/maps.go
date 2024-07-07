package examples

import "fmt"

// MAPs examples
func MapsExamples() {
	var mapp map[string]int			// nil
	mapp = make(map[string]int)		// map[]
	mapp = map[string]int{"two":2}	// map[two: 2]
	mapp["one"] = 1					// map[one:1 two:2]
	delete(mapp, "two")				// map[one:1]
	clear(mapp)						// map[]
	fmt.Println("mapp:", mapp)
}
