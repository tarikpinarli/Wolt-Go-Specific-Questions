# Exercise: Channelling goroutines

This exercise demonstrates how to chain multiple Goroutines together where the **output of one routine becomes the input of another**.

## The Concept
In my answer, I mentioned: *"If the routines need to pass data to each other... I would use unbuffered channels."*

This code implements that concept using a **Pipeline**:
1.  **Source (Generator):** Converts a static slice `[1, 2, 3, 4]` into a stream of data.
2.  **Stage 2 (Square):** Receives an integer, performs a calculation (Square), and passes it down the line.
3.  **Sink (Main):** Consumes the final result.

## Key Mechanics
* **Unidirectional Channels (`<-chan int`):** We use function signatures to enforce direction. `sq` only *reads* from `in` and only *writes* to a new channel. This makes the data flow obvious and safe.
* **Closing Channels:** This is vital in pipelines.
    * When `sliceToChannel` finishes sending data, it calls `close(out)`.
    * This causes the loop `for n := range in` inside the *next* stage to finish.
    * If we forgot to close, the program would crash with a **Deadlock** because the next stage would wait forever for data that never comes.

## 🚀 How to Run
```bash
go run main.go