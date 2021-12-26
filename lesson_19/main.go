package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

/* --------------------------------------------------------------------------------- */
func main() {
	mainTask01()
	mainTask02()
}

/* --------------------------------------------------------------------------------- */
func generateIntArray(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(99)
	}
	return arr
}

func sortAscIntArray(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func sortDescIntArray(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}

/* --------------------------------------------------------------------------------- */
func arrayToString(slice []int) string {
	return strings.Join(strings.Fields(fmt.Sprint(slice)), ",")
}
