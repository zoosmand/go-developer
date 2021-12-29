package main

import (
	"reflect"
	"testing"
)

type ArrayMerger struct {
	arr1 []int
	arr2 []int
}

type TestDataMerger struct {
	name string
	arg  ArrayMerger
	want []int
}

func TestMergeSortedArrays(t *testing.T) {
	tests := []TestDataMerger{
		{
			"first",
			ArrayMerger{
				[]int{7, 13, 15, 17, 20, 23, 78, 88},
				[]int{6, 59, 95, 99},
			},
			[]int{6, 7, 13, 15, 17, 20, 23, 59, 78, 88, 95, 99},
		},
		{
			"second",
			ArrayMerger{
				[]int{6, 59, 95, 99},
				[]int{7, 13, 15, 17, 20, 23, 78, 88},
			},
			[]int{6, 7, 13, 15, 17, 20, 23, 59, 78, 88, 95, 99},
		},
		{
			"third",
			ArrayMerger{
				[]int{1, 2, 3, 4, 5, 6},
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9},
		},
		{
			"fourth",
			ArrayMerger{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{1, 2, 3, 4, 5, 6},
			},
			[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			arrayMerged := make([]int, len(test.arg.arr1)+len(test.arg.arr2))
			mergeSortedArrays(test.arg.arr1, test.arg.arr2, arrayMerged)
			if !reflect.DeepEqual(arrayMerged, test.want) {
				t.Log("func bubbleSortAsc(arr []int)")
				t.Logf("Testing data name: %s", test.name)
				t.Errorf("expected: %v", test.want)
				t.Errorf("got:      %v", arrayMerged)
			}
		})
	}
}
