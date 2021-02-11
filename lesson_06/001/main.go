package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 6.1. Написание простого цикла\n\n")

	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	fmt.Println()

	for i := 0; i < number; i++ {
		fmt.Println(i + 1)
	}

	fmt.Println()
}
