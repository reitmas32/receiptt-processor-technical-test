package tools

import (
	"testing"
)

func TestLenAlphanumerics_AllAlphanumeric(t *testing.T) {
	input := "abc123XYZ456"
	expected := 12
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}

func TestLenAlphanumerics_ExmapleOneTechnicalTest(t *testing.T) {
	input := "Target"
	expected := 6
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}

func TestLenAlphanumerics_ExmapleTwoTechnicalTest(t *testing.T) {
	input := "M&M Corner Market"
	expected := 14
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}

func TestLenAlphanumerics_SomeNonAlphanumeric(t *testing.T) {
	input := "abc 123_ XYZ456"
	expected := 12
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}

func TestLenAlphanumerics_AllNonAlphanumeric(t *testing.T) {
	input := "!@#$%^&*()"
	expected := 0
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}

func TestLenAlphanumerics_EmptyString(t *testing.T) {
	input := ""
	expected := 0
	result := LenAlphanumerics(input)
	if result != expected {
		t.Errorf("For input '%s', expected %d but got %d", input, expected, result)
	}
}
