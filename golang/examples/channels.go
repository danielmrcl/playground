package examples

import (
	"fmt"
	"time"
)

// Channels: are pipes that connect concurrent goroutines.
func ChannelsExamples() {
	// Creating one channel of string messages.
	messages := make(chan string)

	for i := 0; i < 3; i++ {
		// Init an goroutine with anonymous function
		go func() {
			// Do some async shit
			time.Sleep(1 * time.Second)

			if (i == 2) {
				messages <- fmt.Sprintf("Success %d", i)
			}

			log(fmt.Sprintf("%d End of Function", i))
		}()
	}

	// Get the value in the channel
	log("Message received: " + <-messages)

	log("Done!")
}

func log(message string) {
	fmt.Printf("(channels) %s\n", message)
}
