# Exercise: Wrapped Errors (errors.Is)

This exercise demonstrates why we need `errors.Is` when dealing with wrapped errors in Go 1.13+.

## The Problem
When we add context to an error (e.g., "query failed: user not found"), we are "wrapping" the original error inside a new message.
* **Original:** `ErrUserNotFound`
* **Wrapped:** `fmt.Errorf("... %w", ErrUserNotFound)`

Because the wrapped error is a new object, a simple equality check (`err == ErrUserNotFound`) will **fail**.

## The Solution
`errors.Is(err, target)` works like recursively opening a set of boxes. It unwraps the error layer by layer to see if the target error is hiding somewhere inside.

## How to Run
```bash
go run main.go