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

	fmt.Printf("Исходные массивы:\n%v\n%v\n\n", arr4, arr5)

	bubbleSortAsc(arr4)
	bubbleSortAsc(arr5)
	fmt.Printf("Отсортированные массивы:\n%v\n%v\n\n", arr4, arr5)

	mergedArray := make([]int, (len(arr4) + len(arr5)))
	mergeSortedArrays(arr4, arr5, mergedArray)
	fmt.Printf("Объединенный массив:\n%v\n\n", mergedArray)
}

/* --------------------------------------------------------------------------------- */
func mergeSortedArrays(arr1 []int, arr2 []int, arrMerged []int) {
	k := -1
	jj, ii := 0, 0

	for i := ii; i < len(arr1); i++ {
		for j := jj; j < len(arr2); j++ {
			k++
			if arr1[i] < arr2[j] {
				arrMerged[k] = arr1[i]
				ii++
				break
			} else {
				arrMerged[k] = arr2[j]
				jj++
			}
		}
		// fill the rest from arr1
		if len(arr2) == jj {
			for i := ii; i < len(arr1); i++ {
				k++
				arrMerged[k] = arr1[i]
				ii++
			}
		}
	}
	// fill the rest from arr2
	if len(arr2) > jj {
		for j := jj; j < len(arr2); j++ {
			k++
			arrMerged[k] = arr2[j]
		}
	}
}

/* --------------------------------------------------------------------------------- */
// func mergeArrays(args ...[]int) []int {
// 	newArray := []int{}
// 	for _, v := range args {
// 		newArray = append(newArray, v...)
// 	}
// 	return newArray
// }
