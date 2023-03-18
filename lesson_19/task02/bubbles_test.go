package task02

import (
	"reflect"
	"testing"
)

type testDataSorting struct {
	name string
	arg  []int
	want []int
}

func TestBubbleSortAsc(t *testing.T) {
	tests := []testDataSorting{
		{
			"first",
			[]int{16, 877, 22, 81, 345, 69, 59, 83, 29, 776, 46, 920},
			[]int{16, 22, 29, 46, 59, 69, 81, 83, 345, 776, 877, 920},
		},
		{
			"second",
			[]int{552, 842, 467, 4, 81, 0, 78, 512, 12, 9, 88, 134},
			[]int{0, 4, 9, 12, 78, 81, 88, 134, 467, 512, 552, 842},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BubbleSortAsc(test.arg)
			if !reflect.DeepEqual(test.arg, test.want) {
				t.Log("func bubbleSortAsc(arr []int)")
				t.Logf("Testing data name: %s", test.name)
				t.Errorf("expected: %v", test.want)
				t.Errorf("got:      %v", test.arg)
			}
		})
	}
}

func TestBubbleSortDesc(t *testing.T) {
	tests := []testDataSorting{
		{
			"first",
			[]int{16, 877, 22, 81, 345, 69, 59, 83, 29, 776, 46, 920},
			[]int{920, 877, 776, 345, 83, 81, 69, 59, 46, 29, 22, 16},
		},
		{
			"second",
			[]int{552, 842, 467, 4, 81, 0, 78, 512, 12, 9, 88, 134},
			[]int{842, 552, 512, 467, 134, 88, 81, 78, 12, 9, 4, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BubbleSortDesc(test.arg)
			if !reflect.DeepEqual(test.arg, test.want) {
				t.Log("func bubbleSortDesc(arr []int)")
				t.Logf("Testing data name: %s", test.name)
				t.Errorf("expected: %v", test.want)
				t.Errorf("got:      %v", test.arg)
			}
		})
	}
}
