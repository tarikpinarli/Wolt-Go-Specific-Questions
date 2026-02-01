package main

import (
	"errors"
	"fmt"
)

// 1. Define a Custom Struct Error
// This allows to carry specific data
type DatabaseError struct {
	Query     string
	ErrorCode int
}

// Implementing the error interface
func (e *DatabaseError) Error() string {
	return fmt.Sprintf("db error %d: query '%s' failed", e.ErrorCode, e.Query)
}

func main() {
	err := runQuery("SELECT * FROM users")

	if err != nil {
		// We want to access the specific 'ErrorCode' inside the error.
		// We cannot do this with a normal 'err' interface.
		
		var dbErr *DatabaseError

		// The Solution: errors.As
		// It checks: "Does the error err contain a *DatabaseError inside it?"
		// If yes, it fills the dbErr variable with the data.
		if errors.As(err, &dbErr) {
			fmt.Printf("Caught Custom Error! Code: %d, Query: %s\n", dbErr.ErrorCode, dbErr.Query)
		} else {
			fmt.Println("Unknown error")
		}
	}
}

func runQuery(query string) error {
	// Return the custom struct with data
	return &DatabaseError{
		Query:     query,
		ErrorCode: 503, // Service Unavailable
	}
}