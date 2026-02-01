package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Create a dedicated 'done' channel
	stopSignal := make(chan struct{})

	// 2. Start the worker
	fmt.Println("Main: Starting worker...")
	go worker(stopSignal)

	// 3. Let it work for 2 seconds
	time.Sleep(2 * time.Second)

	// 4. BROADCAST THE STOP SIGNAL
	fmt.Println("Main: Closing channel to stop worker...")
	close(stopSignal) // Closing a channel alerts ALL listeners immediately

	time.Sleep(1 * time.Second)
}

func worker(stopCh <-chan struct{}) {
	defer fmt.Println("Worker: Stopped.")

	for {
		select {
		// Check if the channel is closed
		case <-stopCh:
			fmt.Println("Worker: Stop signal received via channel close!")
			return
		default:
			fmt.Println("Worker: Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}