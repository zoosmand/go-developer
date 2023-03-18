package task02

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestReverseIntArray(t *testing.T) {
	testSlice := []int{16, 877, 22, 81, 345, 69, 59, 83, 29, 776, 46, 920}
	checkSlice := []int{920, 46, 776, 29, 83, 59, 69, 345, 81, 22, 877, 16}

	reverseIntArray(testSlice)

	if reflect.DeepEqual(testSlice, checkSlice) {
		return
	} else {
		t.Errorf("Expected %s, got %s",
			arrayToString(checkSlice),
			arrayToString(testSlice),
		)
	}
}

func arrayToString(slice []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(slice)), ",")
}
