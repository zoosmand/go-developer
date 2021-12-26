package main

import (
	"reflect"
	"testing"
)

func TestSortDescIntArray(t *testing.T) {
	testSlice := []int{16, 877, 22, 81, 345, 69, 59, 83, 29, 776, 46, 920}
	checkSlice := []int{920, 877, 776, 345, 83, 81, 69, 59, 46, 29, 22, 16}

	sortDescIntArray(testSlice)

	if reflect.DeepEqual(testSlice, checkSlice) {
		return
	} else {
		t.Errorf("Expected %s, got %s",
			arrayToString(checkSlice),
			arrayToString(testSlice),
		)
	}
}
