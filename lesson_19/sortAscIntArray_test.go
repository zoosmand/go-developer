package main

import (
	"reflect"
	"testing"
)

func TestSortAscIntArray(t *testing.T) {
	testArray := [12]int{16, 877, 22, 81, 345, 69, 59, 83, 29, 776, 46, 920}
	checkArray := [12]int{16, 22, 29, 46, 59, 69, 81, 83, 345, 776, 877, 920}

	sortAscIntArray(testArray[0:])

	if reflect.DeepEqual(testArray, checkArray) {
		return
	} else {
		t.Errorf("Expected %s, got %s",
			arrayToString(checkArray[0:]),
			arrayToString(testArray[0:]),
		)
	}
}
