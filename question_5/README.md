# Question 5: Test Coverage & Table-Driven Tests

This assignment demonstrates **Go's Table-Driven Test** pattern to handle edge cases effectively.

## The Scenario
I implemented a simple `CalculateGrade` function. Initially, simple manual tests missed critical bugs such as **negative inputs** or **scores above 100**.

## The Solution
I restructured the tests using a **Table-Driven** approach. This allowed me to:
1.  Define **Happy Paths** (Valid inputs).
2.  Define **Edge Cases** (Negative numbers, boundaries).
3.  Use `t.Run` for isolated **Subtests**.

## How to Run

To run the tests with verbose output:
```bash
go test -v