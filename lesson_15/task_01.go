package main

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func mainTask01() {
	h := "Задача 15.1. Подсчёт чётных и нечётных чисел в массиве."
	fmt.Print(usecons.Header(&h))

	const arrLen int = 10
	var digits [arrLen]int

	// ----------------------------------------------------
	var consoleInput int
	invite := fmt.Sprintf("Введите %v чисел. (последовательно или одной строкой через запятую)", arrLen)
	fmt.Println(invite)

	// ----------------------------------------------------
	i := 0
	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			digits[i] = consoleInput
			i++
			if i >= arrLen {
				break
			}
		}
	}
	// send slice of the orig array "digits" full of length as a paramater
	evenCount, oddCount := oddEvenCount(digits[0:])

	fmt.Printf("\nЧетные: %v, нечетные: %v", evenCount, oddCount)
}

func oddEvenCount(arr []int) (oddCount int, evenCount int) {
	for _, v := range arr {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return oddCount, evenCount
}
