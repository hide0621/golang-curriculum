package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) returned %d, expected 3.", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(2, 1)
	expected := 1
	if result != expected {
		t.Errorf("Subtract(2, 1) returned %d, expected 1.", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(6, 3)
	if err != nil {
		t.Errorf("Divide(6, 3) failed. Expected no error, got %v", err.Error())
	}
	expected := 2
	if result != expected {
		t.Errorf("Divide(6, 3) returned %d, expected 2.", result)
	}
}
