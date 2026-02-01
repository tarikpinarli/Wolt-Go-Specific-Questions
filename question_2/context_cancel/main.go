package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. Creating the Context (the signal receiver) and Cancel function (the stop button)
	ctx, cancel := context.WithCancel(context.Background())

	// 2. Start the background worker
	fmt.Println("Main: Starting worker...")
	go worker(ctx)

	// 3. Letting the go routine to work for 2 seconds
	time.Sleep(2 * time.Second)

	// 4. PRESS THE STOP BUTTON
	fmt.Println("Main: Time's up! Cancelling context...")
	cancel() // This sends the signal to <-ctx.Done()

	// 5. Waiting a moment to see the worker print its cleanup message
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting.")
}

func worker(ctx context.Context) {
	// Verify exit. This defer will make the print the last thing this routine does
	defer fmt.Println("Worker: Stopped (cleanup complete).")

	for {
		select {
		// The "Signal Receiver"
		case <-ctx.Done():
			fmt.Println("Worker: I received the cancel signal! Shutting down...")
			return // Exit
		default:
			// Default work
			fmt.Println("Worker: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}