package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel (The Bridge)
	// It has no storage capacity.
	dataBridge := make(chan string)

	// Goroutine 1: The Sender
	go func() {
		fmt.Println("Sender: I have data, walking to the bridge...")
		time.Sleep(1 * time.Second) // Simulate some work
		
		// Sending data. This BLOCKS until someone is ready to receive.
		// It puts the data directly on the bridge, not in a box.
		dataBridge <- "Secret Payload" 
		
		fmt.Println("Sender: Data received! I can leave now.")
	}()

	// Main Goroutine: The Receiver
	fmt.Println("Receiver: Waiting at the bridge...")
	
	// Receiving data. This BLOCKS until data is sent.
	message := <-dataBridge
	
	fmt.Println("Receiver: Got the message:", message)
}