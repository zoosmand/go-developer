package main

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

/* --------------------------------------------------------------------------------- */
func mainTask01() {
	h := "Задача 19.1. Слияние отсортированных массивов."
	fmt.Print(usecons.Header(&h))

	arr4 := generateIntArray(4)
	arr5 := generateIntArray(5)

	fmt.Printf("Исходные массивы:\n%s\n%s\n\n", arrayToString(arr4), arrayToString(arr5))

	sortAscIntArray(arr4)
	sortAscIntArray(arr5)
	fmt.Printf("Отсортированные массивы:\n%s\n%s\n\n", arrayToString(arr4), arrayToString(arr5))

	mergedArray := mergeArrays(arr4, arr5)
	fmt.Printf("Объединенный массив:\n%s\n\n", arrayToString(mergedArray))
}

/* --------------------------------------------------------------------------------- */
func mergeArrays(args ...[]int) []int {
	newArray := []int{}
	for _, v := range args {
		newArray = append(newArray, v...)
	}
	return newArray
}
