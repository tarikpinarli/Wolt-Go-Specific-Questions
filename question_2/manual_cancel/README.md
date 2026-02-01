# Exercise: Manual Cancellation (The "Old Way")

This exercise demonstrates how cancellation works under the hood by using a raw channel. This was the common pattern in Go before the `context` package became standard.

## The Concept
A channel in Go can be used as a broadcast mechanism. When you **close** a channel, every Goroutine waiting on that channel receives a signal immediately.

* We create a `stopSignal` channel.
* The worker listens to it.
* When we want to stop the worker, we `close(stopSignal)`.

## Comparison to Context
While this works for simple cases, `context` is superior because:
1.  **Hierarchy:** Contexts can be nested (e.g., a Database timeout inside an HTTP request).
2.  **Values:** Contexts can carry data (Trace IDs, User IDs).
3.  **Standardization:** All modern Go libraries expect a `context` object, not a custom channel.

## How to Run
```bash
go run main.go