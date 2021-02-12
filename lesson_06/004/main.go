package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.4. Предыдущий урок на if\n\n")

	var first int
	var second int
	var third int

	for true {
		if third >= 1000 {
			break
		}
		third++

		if second >= 100 {
			continue
		}
		second++

		if first >= 10 {
			continue
		}
		first++
	}

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
