package examples

import (
	"fmt"
	"time"
)

// Buffered Channels: are channels that accept a limited number of
// values without a corresponding receiver.
func BufferedChannelsExamples() {
	log := logger{"buffered channels"}.log

	// Creating one channel of 3 string messages.
	messages := make(chan string, 3)

	for i := 0; i < 3; i++ {
		// Init an goroutine with anonymous function
		go func() {
			// Do some async shit
			time.Sleep(1 * time.Second)

			messages <- fmt.Sprintf("%d End of Function", i)
		}()
	}

	// Get the value in the channel
	log("Message received: " + <-messages)
	log("Message received: " + <-messages)
	log("Message received: " + <-messages)

	log("Done!")
}
