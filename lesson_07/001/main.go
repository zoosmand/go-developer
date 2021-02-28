package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.1. Зеркальные билеты\n\n")
	var count int

	for i := 100000; i <= 999999; i++ {
		head := i / 1000
		tail := i % 1000

		if head == (tail/100 + tail%10*100 + tail/10%10*10) {
			count++
			// fmt.Println(i)
		}

		// var newTail int
		// z := 0
		// for k := 2; k >= 0; k-- {
		// 	newTail += int(math.Pow10(z)) * (tail / int(math.Pow10(k)) % 10)
		// 	z++
		// }

		// if head == newTail {
		// 	count++
		// 	// fmt.Println(i)
		// }
	}

	fmt.Println(count)
}
