package task02

import (
	"fmt"

	"lesson_19/zoosmand/requisites"

	"github.com/zoosmand/usecons/v3"
)

/* --------------------------------------------------------------------------------- */
func BubbleSorting() {
	h := "Задача 19.1. Сортировка пузырьком."
	fmt.Print(usecons.Header(&h))

	arr100 := requisites.GenerateIntArray(100)
	fmt.Printf("Исходный массив:\n%v\n\n", arr100)

	BubbleSortAsc(arr100)
	fmt.Printf("Массив, отсортированный в прямом порядке пузырьками:\n%v\n\n", arr100)

	BubbleSortDesc(arr100)
	fmt.Printf("Массив, отсортированный в обратном порядке пузырьками:\n%v\n\n", arr100)
}

/* --------------------------------------------------------------------------------- */
func BubbleSortAsc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

/* --------------------------------------------------------------------------------- */
func BubbleSortDesc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
