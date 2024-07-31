package examples

import (
	"fmt"
	"sync"
	"time"
)

// Examples of Goroutines and WaitGroup.
// - *Goroutine* is a lightwaight thread of execution.
// - *WaitGroup* is used to group multiples goroutines
//   and wait the finish of all to continue.
func GoroutinesExamples() {
	fmt.Println("(goroutines) Running all functions...")

	// Group to wait all routines
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		// Increment the num of routines
		wg.Add(1)

		// Init an goroutine with anonymous function
		go func() {
			// Do some async shit
			time.Sleep(1 * time.Second)
			fmt.Printf("(goroutines) %d End of Function\n", i)

			// Marks the routine as Done
			wg.Done()
		}()
	}

	// Wait all routines in the group
	wg.Wait()

	fmt.Println("(goroutines) Done!")
}
