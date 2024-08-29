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
		// Channel as write-only (chan<-)
		go func(msg chan<- string) {
			// Do some async shit
			time.Sleep(1 * time.Second)

			msg <- fmt.Sprintf("%d End of Function", i)
		}(messages)
	}

	// Channel as read-only (<-chan)
	func(msg <-chan string) {
		// Get the value in the channel
		log("Message received: " + <-msg)
		log("Message received: " + <-msg)
		log("Message received: " + <-msg)
	}(messages)

	log("Done!")
}
