# Wolt Internship 2026 - Technical Demonstrations

## Purpose of this Repository
This repository serves as a **technical supplement** to my application for the Wolt Summer 2026 Internship (Backend Go Track).

In my application form, I provided written answers. I created these code samples to strictly demonstrate those concepts in practice and to validate the logic described in my answers.

## Repository Structure & Application Mapping

### 1. Concurrency Orchestration (Question 1)
*Located in `/question_1`*

My answer discussed the difference between independent tasks ("Fork-Join") and data-dependent tasks ("CSP").
* **[question_1/waitgroup/]:** Demonstrates synchronization of independent workers (The "Coffee Shop" example).
* **[question_1/channels/]:** A more advanced implementation showing how to chain Goroutines where the output of one stage is the input of the next.

### 2. Goroutine Cancellation (Question 2)
*Located in `/question_2`*

My answer emphasized the importance of resource management and the shift from manual signaling to the `context` package.
* **[question_2/context_cancel/]:** The modern, standard approach using `context.WithCancel` to signal multiple workers to stop.
* **[question_2/manual_channel/]:** The "Old Way" (broadcasting via channel closure), demonstrating my understanding of what `context` actually does under the hood.

### 3. Error Handling (Question 3)
*Located in `/question_3`*

My answer differentiated between Wrapped and Custom Struct Errors.
* **[question_3/wrapped_errors/]:** Demonstrates why `err == target` fails with wrapped errors and how `errors.Is` solves it.
* **[question_3/struct_errors/]:** Demonstrates using `errors.As` to safely cast generic errors back to custom structs to access specific data fields.

### 4. Test Coverage (Question 5)
*Located in `/question_5`*

My answer discussed using **Table-Driven Tests** to catch bugs and handle edge cases systematically.
* **[question_5/]:** Contains a `CalculateGrade` function with a comprehensive test suite using `t.Run`. It demonstrates how I handle boundary conditions (scores of 0 or 100) and invalid inputs (negative scores) using Go's standard testing patterns.

## How to Run
Each folder is a self-contained Go module. You can run any example by navigating to the folder and executing `go run main.go`.

For the testing exercise, run:
```bash
cd testing
go test -v