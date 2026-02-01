# Exercise: Custom Struct Errors (errors.As)

This exercise demonstrates how to access specific data fields (like error codes) from a custom error struct using `errors.As`.

## The Concept
Sometimes an error needs to carry data, not just a message. We use a custom struct (`DatabaseError`) to hold fields like `ErrorCode` and `Query`.

However, Go functions usually return the generic `error` interface. This hides the struct fields. We need a way to "cast" the generic error back to our specific struct to read the data.

## How It Works
* **The Struct:** We define `DatabaseError` and implement the `Error()` method to satisfy the interface.
* **The Cast (`errors.As`):**
    * We declare an empty pointer: `var dbErr *DatabaseError`.
    * We call `errors.As(err, &dbErr)`.
    * This function checks if `err` contains a `*DatabaseError`. If it does, it points `dbErr` to that data.

## How to Run
```bash
go run main.go