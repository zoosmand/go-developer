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
	
	for i := 0; i < size; i++ {
		for k := 0; k < size; k++ {
			if (i^k)&1 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("---------------------")
}
