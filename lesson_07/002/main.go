package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.2. Шахматная доска\n\n")

	var size int
	fmt.Println("Введите размерность доски: ")
	fmt.Scan(&size)

	fmt.Println("---------------------")

	field := '*'
	var toggleField = func() {
 		if field == ' ' {
			field = '*'
		} else {
			field = ' '
		}
	}

	for i := 0; i < size; i++ {
		var line []rune
		for k := 0; k < size; k++ {
			toggleField()
			line = append(line, field)
		}
		// Depending on parity skip the next field
		if size%2 == 0 {
			toggleField()
		}
		fmt.Println(string(line))
	}

	fmt.Println("---------------------")
}
