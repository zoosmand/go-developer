package main

import "testing"

func TestOddEvenCount(t *testing.T) {
	testArray := [10]int{1, 2, 3, 41, 5, 6, 7, 8, 9, 0}
	odd, even := oddEvenCount(testArray[0:])
	expectedOdd, expectedEven := 6, 4

	if odd != expectedOdd || even != expectedEven {
		t.Errorf(
			"Expected odd - %d, got %d, expected even - %d, got %d",
			expectedOdd, odd, expectedEven, even)
	}
}
