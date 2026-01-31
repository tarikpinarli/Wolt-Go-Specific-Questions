package main

import (
	"errors"
	"fmt"
)

// CalculateGrade determines the letter grade based on the score (0-100).
// It returns an error for invalid inputs (negative numbers or >100).
func CalculateGrade(score int) (string, error) {
	// Edge Case 1: Negative numbers are not allowed
	if score < 0 {
		return "", errors.New("score cannot be negative")
	}
	// Edge Case 2: Scores higher than 100 are not allowed
	if score > 100 {
		return "", errors.New("score cannot be higher than 100")
	}

	// Standard logic
	if score >= 90 {
		return "A", nil
	} else if score >= 80 {
		return "B", nil
	} else if score >= 70 {
		return "C", nil
	} else if score >= 50 {
		return "Pass", nil
	}

	return "Fail", nil
}

// main function to allow manual running if needed
func main() {
	grade, err := CalculateGrade(95)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grade)
	}
}