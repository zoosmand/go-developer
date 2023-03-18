package requisites

import (
	"math/rand"
)

/* --------------------------------------------------------------------------------- */
func GenerateIntArray(length int) []int {
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(99)
	}
	return arr
}
