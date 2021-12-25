package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestReverseIntArray(t *testing.T) {
	const arrLength = 100
	var testArray [arrLength]int
	var resultArray [arrLength]int

	for i := 0; i < arrLength; i++ {
		testArray[i] = rand.Int()
	}

	// resultArray := testArray
	copy(resultArray[0:], testArray[0:])

	reverseIntArray(resultArray[0:])

	if reflect.DeepEqual(testArray, resultArray) {
		// if testArray == resultArray {
		t.Errorf("Expected %s, got %s",
			strings.Join(strings.Fields(fmt.Sprint(testArray[0:])), ","),
			strings.Join(strings.Fields(fmt.Sprint(resultArray[0:])), ","),
		)
	}
}
