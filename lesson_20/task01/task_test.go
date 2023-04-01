package task01

import (
	"reflect"
	"testing"

	"github.com/zoosmand/usecons/v3"
)

// testing the levels from 2 to 5
func TestTransposeMatrix(t *testing.T) {
	var m, m2, m3 [][]int

	for i := 2; i < 6; i++ {
		m = usecons.GenerateIntSquareMatrix(i)
		m2 = transposeMatrix(m)
		m3 = transposeMatrix(m2)

		if !reflect.DeepEqual(m, m3) {
			t.Errorf("The doubly transposed matrix of level %v should be equal to the initial one.\n", i)
		}
	}
}

// testing the levels from 2 to 5
func TestCalculateDeterminant(t *testing.T) {
	var m1, m2 [][]int
	var d1, d2 int

	for i := 2; i < 6; i++ {
		m1 = usecons.GenerateIntSquareMatrix(i)
		m2 = transposeMatrix(m1)

		d1 = calculateDeterminant(m1)
		d2 = calculateDeterminant(m2)

		if d1 != d2 {
			t.Errorf("The determinant '%v'  of the matrix of level '%v' should be equal to the determinant '%v' of the same transponsed matrix.\n", d1, i, d2)
		}
	}
}
