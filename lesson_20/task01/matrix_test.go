package task01

import (
	"reflect"
	"testing"

	"github.com/zoosmand/usecons/v3"
)

func TestTransposeMatrixLevel2(t *testing.T) {
	level := 2
	m := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m)
	m3 := transposeMatrix(m2)

	if !reflect.DeepEqual(m, m3) {
		t.Errorf("The doubly transposed matrix of level %v should be equal to the initial one.\n", level)
	}
}

func TestTransposeMatrixLevel3(t *testing.T) {
	level := 3
	m := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m)
	m3 := transposeMatrix(m2)

	if !reflect.DeepEqual(m, m3) {
		t.Errorf("The doubly transposed matrix of level %v should be equal to the initial one.\n", level)
	}
}

func TestTransposeMatrixLevel4(t *testing.T) {
	level := 4
	m := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m)
	m3 := transposeMatrix(m2)

	if !reflect.DeepEqual(m, m3) {
		t.Errorf("The doubly transposed matrix of level %v should be equal to the initial one.\n", level)
	}
}

func TestTransposeMatrixLevel5(t *testing.T) {
	level := 5
	m := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m)
	m3 := transposeMatrix(m2)

	if !reflect.DeepEqual(m, m3) {
		t.Errorf("The doubly transposed matrix of level %v should be equal to the initial one.\n", level)
	}
}

func TestCalculateDeterminantLevel2(t *testing.T) {
	level := 2
	m1 := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m1)

	d1 := calculateDeterminant(m1)
	d2 := calculateDeterminant(m2)

	if d1 != d2 {
		t.Errorf("The determinant '%v'  of the matrix of level '%v' should be equal to the determinant '%v' of the same transponsed matrix.\n", d1, level, d2)
	}
}

func TestCalculateDeterminantLevel3(t *testing.T) {
	level := 3
	m1 := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m1)

	d1 := calculateDeterminant(m1)
	d2 := calculateDeterminant(m2)

	if d1 != d2 {
		t.Errorf("The determinant '%v'  of the matrix of level '%v' should be equal to the determinant '%v' of the same transponsed matrix.\n", d1, level, d2)
	}
}

func TestCalculateDeterminantLevel4(t *testing.T) {
	level := 4
	m1 := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m1)

	d1 := calculateDeterminant(m1)
	d2 := calculateDeterminant(m2)

	if d1 != d2 {
		t.Errorf("The determinant '%v'  of the matrix of level '%v' should be equal to the determinant '%v' of the same transponsed matrix.\n", d1, level, d2)
	}
}

func TestCalculateDeterminantLevel5(t *testing.T) {
	level := 5
	m1 := usecons.GenerateIntSquareMatrix(level)
	m2 := transposeMatrix(m1)

	d1 := calculateDeterminant(m1)
	d2 := calculateDeterminant(m2)

	if d1 != d2 {
		t.Errorf("The determinant '%v'  of the matrix of level '%v' should be equal to the determinant '%v' of the same transponsed matrix.\n", d1, level, d2)
	}
}
