package task02

import (
	"fmt"
	"lesson_20/zoosmand/task01"

	"github.com/zoosmand/usecons/v3"
)

func MultiplyMatrices() {
	h := "Task 20.2 Multiplying two matrices"
	fmt.Println(usecons.Header(&h))

	h1, w1 := 3, 5
	m1 := usecons.GenerateIntMatrix(h1, w1)

	h2, w2 := 5, 4
	m2 := usecons.GenerateIntMatrix(h2, w2)

	task01.PrintMatix(&m1)
	fmt.Println()
	task01.PrintMatix(&m2)

	rm := multiplyMatrices(m1, m2)

	fmt.Println()
	task01.PrintMatix(&rm)
	fmt.Println()
}

func multiplyMatrices(m1 [][]int, m2 [][]int) [][]int {
	if len(m1[0]) != len(m2) {
		panic("The matrices are not consistent.")
	}

	rm := make([][]int, len(m1))

	for i := range m1 {
		rm[i] = make([]int, len(m2[0]))
		for j := range m2[0] {
			for r := range m2 {
				rm[i][j] += m1[i][r] * m2[r][j]
			}
		}
	}
	return rm
}
