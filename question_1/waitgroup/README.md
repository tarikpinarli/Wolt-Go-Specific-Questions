# Exercise: Sync.WaitGroup (The Coffee Shop)

This exercise demonstrates the **Fork-Join** concurrency model in Go using `sync.WaitGroup`. It simulates a coffee shop where multiple orders ("latte" and "americano") are processed simultaneously by independent workers.

## The Concept
When we have multiple tasks that are completely independent of each other (like two different baristas making two different drinks), we don't need them to talk to each other. We just need to know when **all** of them are finished.

We use `sync.WaitGroup` to orchestrate this:
1.  **Add**: We tell the counter how many tasks we are starting.
2.  **Done**: Each task decrements the counter when it finishes.
3.  **Wait**: The main function blocks (sleeps) until the counter reaches zero.

## How It Works
* **Structs:** We define an `Order` struct to hold the ID and item name.
* **Methods:** The `Process()` method simulates work using `time.Sleep`.
* **Concurrency:** We launch two anonymous Goroutines. Each calls `wg.Done()` using `defer` to ensure the counter is decremented even if the function crashes.

## How to Run
```bash
go run main.go