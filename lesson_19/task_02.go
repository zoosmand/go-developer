package main

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

/* --------------------------------------------------------------------------------- */
func mainTask02() {
	h := "Задача 19.1. Сортировка пузырьком."
	fmt.Print(usecons.Header(&h))

	arr100 := generateIntArray(100)
	fmt.Printf("Исходный массив:\n%v\n\n", arr100)

	bubbleSortAsc(arr100)
	fmt.Printf("Массив, отсортированный в прямом порядке пузырьками:\n%v\n\n", arr100)

	bubbleSortDesc(arr100)
	fmt.Printf("Массив, отсортированный в обратном порядке пузырьками:\n%v\n\n", arr100)
}

/* --------------------------------------------------------------------------------- */
func bubbleSortAsc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

/* --------------------------------------------------------------------------------- */
func bubbleSortDesc(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
