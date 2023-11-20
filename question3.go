package main

import (
	"fmt"
	"time"
)

func question3() {
	// Create channel for making signal when the work is done
	done := make(chan bool)

	go func() {
		// Wait for some operation
		time.Sleep(3 * time.Second)

		// When the work is done the channel set with true
		done <- true
	}()

	select {
	// Waiting for complete the work
	case <-done:
		// The goroutine completed before the timeout
		fmt.Println("Goroutine completed successfully.")
	case <-time.After(2 * time.Second):
		// Timeout occurred
		fmt.Println("Timeout reached. Goroutine did not complete in time.")
	}
}
