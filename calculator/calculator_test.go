package calculator

import (
	"testing"
)

func TestSquareNumbers(t *testing.T) {
	first := 2
	second := 3
	channel := SquareNumbers(first, second)

	firstGot := <-channel
	firstWant := first * first
	if  firstGot != firstWant {
		t.Errorf("SquareNumbers() = %q, want %q", firstGot, firstWant)
	}

	secondGot := <-channel
	secondWant := first * first
	if  firstGot != firstWant {
		t.Errorf("SquareNumbers() = %q, want %q", secondGot, secondWant)
	}
}