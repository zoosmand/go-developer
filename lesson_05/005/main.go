package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.5. Определение максимальных процентов")
	fmt.Println()

	var percents []int

	for i := 0; i < 3; i++ {
		var percent int
		fmt.Println("Введите процентную ставку: ")
		fmt.Scan(&percent)
		percents = append(percents, percent)
	}

	sort.Ints(percents)

	fmt.Println("Наибльшие процентные ставки:", percents[1:3])
	fmt.Println()
}
