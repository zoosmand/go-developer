package task01

import "testing"

func TestOddEvenCount(t *testing.T) {
	testArray := [12]int{1, 2, 3, 1, 5, 6, 7, 8, 9, 6, 6, 2}
	odd, even := oddEvenCounter(testArray[0:])
	expectedOdd, expectedEven := 6, 6

	if odd != expectedOdd || even != expectedEven {
		t.Errorf(
			"Expected odd - %d, got %d, expected even - %d, got %d",
			expectedOdd, odd, expectedEven, even)
	}
}
