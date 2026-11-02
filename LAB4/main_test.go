package main

import "testing"

func TestDivide(t *testing.T) {
	// Valid input
	result, err := divide(10, 2)
	if err != nil || result != 5 {
		t.Errorf("Expected 5, got %v, error: %v", result, err)
	}

	// Division by zero
	_, err = divide(10, 0)
	if err == nil {
		t.Error("Expected division by zero error")
	}
}

func TestSquareRoot(t *testing.T) {
	// Valid input
	result, err := squareRoot(25)
	if err != nil || result != 5 {
		t.Errorf("Expected 5, got %v, error: %v", result, err)
	}

	// Negative number
	_, err = squareRoot(-9)
	if err == nil {
		t.Error("Expected negative square root error")
	}
}

func TestCheckInput(t *testing.T) {
	// Valid input
	text, err := checkInput("Hello")
	if err != nil || text != "Hello" {
		t.Errorf("Expected Hello, got %q, error: %v", text, err)
	}

	// Empty input
	_, err = checkInput("")
	if err == nil {
		t.Error("Expected empty input error")
	}
}
