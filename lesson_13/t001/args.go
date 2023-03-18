package t001

import (
	"fmt"
	"math"

	"github.com/zoosmand/usecons/v3"
)

func Arguments() {
	h := "Задача 13.1. Функция, принимающая аргументы."
	fmt.Print(usecons.Header(&h))

	fmt.Println(evenSum(234, 568))
}

func evenSum(a, b int) int {
	fa := float64(a)
	fb := float64(b)

	var result int64

	for i := math.Min(fa, fb); i < math.Max(fa, fb); i++ {
		ii := int64(i)
		if ii%2 == 0 {
			result += ii
		}
	}
	return int(result)
}
