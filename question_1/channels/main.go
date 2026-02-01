package main

import (
	"fmt"
)

// Stage 1: The Source (Generator)
// Converts a slice of integers into a stream of data via a channel.
func sliceToChannel(n []int) <-chan int {
	out := make(chan int)

	go func() {
		// Iterate over the slice and send items one by one
		// Herre the tricky part is that: When we range over a channel and if the channel is not closed
		// the goroutine will be blocked at this moment.
		for _, num := range n {
			out <- num
		}
		// Close the channel to signal "no more data"
		// If we don't close, the next stage will wait forever (deadlock).
		close(out)
	}()
	return out
}

// Stage 2: The Transformer (Square)
// Receives data from 'in', processes it, and sends it to 'out'.
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// 'range in' keeps reading until 'in' is closed by the previous stage
		for n := range in {
			out <- n * n
		}
		// Close the output channel when we are done processing
		close(out)
	}()
	return out
}

func main() {
	// 0. The raw data
	nums := []int{1, 2, 3, 4}

	// 1. Set up the Pipeline
	// slice -> [sliceChannel] -> sq -> [squareChannel] -> main
	sliceChannel := sliceToChannel(nums)
	squareChannel := sq(sliceChannel)

	// Stage 3: The Sink (Consumer)
	// Read from the final channel until it is closed
	for n := range squareChannel {
		fmt.Printf("%d\n", n)
	}
}