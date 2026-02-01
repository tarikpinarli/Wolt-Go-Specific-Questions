# Exercise: Cancellation with Context

This exercise demonstrates the **modern standard** for handling cancellation in Go: the `context` package. It shows how to stop a background Goroutine cleanly without leaving resources hanging.

## The Concept
In long-running applications (like a server handling 10,000 requests), we cannot afford to let background tasks run forever if the user disconnects. We need a "Stop Button."

* **The Signal Receiver:** `ctx.Done()`
    * Inside the worker, we constantly check this channel. If it's closed, it means "Stop working now."
* **The Stop Button:** `cancel()`
    * Calling this function (returned by `context.WithCancel`) instantly closes the `Done` channel, broadcasting the signal to the worker.

## How It Works
1.  **Setup:** We create a `context` in `main`.
2.  **Pass:** We pass this context into the `worker` Goroutine.
3.  **Listen:** The worker uses a `select` statement. If `<-ctx.Done()` triggers, it returns immediately.
4.  **Verify:** We use `defer` to print a message when the worker function actually exits, proving that the Goroutine is destroyed and memory is freed.

## How to Run
```bash
go run main.go