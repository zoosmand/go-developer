package task01

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/zoosmand/usecons/v3"
)

func Determinant() {
	h := "Task 20.1 The determinant"
	fmt.Println(usecons.Header(&h))

	var printResult = func(l int) {

		m := usecons.GenerateIntSquareMatrix(l)

		printMatix(&m)
		d := calculateDeterminant(m)

		fmt.Printf("\n---\nThe determinant of the matrix is %v\n---\n", d)
	}

	// ---
	var consoleInput int
	invite := "Input a positive integer value meaning a level of a square matrix\n\n"
	fmt.Print(invite)

	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Print("Wrong input!", invite)
		} else {
			printResult(consoleInput)
			break
		}
	}

}

func calculateDeterminant(m [][]int) int {
	det := 0
	if len(m) == 0 {
		log.Fatal("\n---\nThe martix dimension has to be at least 2.\n")
		os.Exit(-1)
	}
	for i := range m[0] {
		if len(m[0]) > 2 {
			r := reduceMatrix(m, 1, uint(i+1))
			det += int(math.Pow(-1, float64(i))) * calculateDeterminant(r) * m[0][i]
		} else {
			return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
		}
	}
	return det
}

func printMatix(m *[][]int) {
	mm := *m
	for i := range mm {
		fmt.Println(mm[i])
	}
}

func reduceMatrix(m [][]int, row uint, col uint) [][]int {
	if row <= 0 || col <= 0 || row > uint(len(m)) || col > uint(len(m)) {
		log.Fatal("\n---\nThe martix dimension has to be at least 2.\n")
		os.Exit(-1)
	}
	res := make([][]int, len(m))
	copy(res, m)
	for i := range m {
		res[i] = make([]int, len(m[i]))
		copy(res[i], m[i])
	}

	res = append(res[:int(row)-1], res[int(row):]...)

	for i := range res {
		res[i] = append(res[i][:int(col)-1], res[i][int(col):]...)
	}

	return res
}

func transposeMatrix(m [][]int) [][]int {
	s := len(m[0])
	var t [][]int
	for i := range m[0] {
		t = append(t, make([]int, s))
		for j := range m[0] {
			t[i][j] = m[j][i]
		}
	}
	return t
}
