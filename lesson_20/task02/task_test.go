package task02

import (
	"reflect"
	"testing"
)

func TestMultiplyMatrix(t *testing.T) {
	m1 := [][][]int{
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
		},
		{
			{77, 17, 420, 22, 4},
			{608, 92, 65, 773, 0},
			{554, 21, 3, 717, 83},
		},
	}

	m2 := [][][]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
			{17, 18, 19, 20},
		},
		{
			{83, 215, 99, 62},
			{45, 351, 119, 740},
			{6, 35, 185, 23},
			{894, 5, 29, 87},
			{533, 93, 883, 23},
		},
	}

	want := [][][]int{
		{
			{175, 190, 205, 220},
			{400, 440, 480, 520},
			{625, 690, 755, 820},
		},
		{
			{31476, 37704, 91516, 29020},
			{746056, 169152, 105582, 174522},
			{732182, 137890, 151982, 114245},
		},
	}
	if len(m1) != len(m2) || len(m1) != len(want) || len(m2) != len(want) {
		t.Error("Testing data is not ready.\n")
	}

	for i := range m1 {
		if !reflect.DeepEqual(multiplyMatrices(m1[i], m2[i]), want[i]) {
			t.Error("multiplyMatrix(m1, m2) return incorrect result.\n")
		}
	}
}
