package Day2

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestFilterOddNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	actual := FilterOddNumbers(numbers)
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestSquareNumbersInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	actual := SquareNumbers(numbers)
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], actual[i])
		}
	}
}

func TestSquareNumbersFloat(t *testing.T) {
	numbers := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.}
	expected := []float64{0.01, 0.04, 0.09, 0.16, 0.25, 0.36, 0.49, 0.64, 0.81, 1.}
	actual := SquareNumbers(numbers)
	for i := range actual {
		if !almostEqual(actual[i], expected[i]) {
			t.Errorf("Expected %v, got %v", expected[i], actual[i])
		}
	}
}

func TestCountChars(t *testing.T) {
	words := []string{"apple", "banana", "cherry"}
	expected := []int{5, 6, 6}
	actual := CountChars(words)
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	}
}
