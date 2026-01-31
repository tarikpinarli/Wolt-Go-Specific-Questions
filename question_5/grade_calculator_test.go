package main

import "testing"

func TestCalculateGrade(t *testing.T) {
	// 1. Defining the Table (Test Scenarios)
	// We use a struct slice to define inputs and expected outputs clearly.
	tests := []struct {
		name      string // Description of the test case
		input     int    // Input score
		wantGrade string // Expected result
		wantErr   bool   // Do we expect an error?
	}{
		// Happy Paths (Valid Inputs)
		{"High Score (A)", 95, "A", false},
		{"Mid Score (B)", 85, "B", false},
		{"Boundary Pass", 50, "Pass", false},
		{"Fail Score", 40, "Fail", false},

		// Edge Cases (Invalid Inputs - Catching Bugs)
		{"Negative Input", -10, "", true},     // Should return error
		{"Over Limit", 105, "", true},         // Should return error
		{"Boundary Zero", 0, "Fail", false},   // Valid boundary
		{"Boundary Hundred", 100, "A", false}, // Valid boundary
	}

	// 2. The Loop (Running Subtests)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrade, err := CalculateGrade(tt.input)

			// Error Assertion: Check if error presence matches expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateGrade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Value Assertion: Check if the grade matches
			if gotGrade != tt.wantGrade {
				t.Errorf("CalculateGrade() = %v, want %v", gotGrade, tt.wantGrade)
			}
		})
	}
}