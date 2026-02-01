package main

import (
	"errors"
	"fmt"
)

// 1. Define a Sentinel Error (The "Root Cause")
var ErrUserNotFound = errors.New("user not found")

func main() {
	err := findUser("john_doe")

	if err != nil {
		fmt.Println("Error occurred:", err)

		// !!! The Problem: Simple equality check FAILS
		// Because the error is inside a "box" (wrapped), it is not equal to the original.
		if err == ErrUserNotFound {
			fmt.Println("equality check with direct error: Matched")
		} else {
			fmt.Println("equality check with direct error: Failed (Because the error is wrapped)")
		}

		// The Solution: errors.Is
		// It intelligently looks inside the chain to find the root cause.
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("errors.Is: Matched! Handling 404 (user not found)")
		}
	}
}

func findUser(username string) error {
	// Simulating a failure
	// We wrap the error to add context ("database query failed")
	// The %w verb puts ErrUserNotFound inside that database query failed "box"
	return fmt.Errorf("database query failed: %w", ErrUserNotFound)
}