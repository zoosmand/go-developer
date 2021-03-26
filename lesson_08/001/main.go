package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 8.1.\n\n")

	var month int
	fmt.Println("Введите номер месяца от 1 до 12: ")

	for {
		_, _ = fmt.Scan(&month)

		if month >= 1 && month <= 12 {
			break
		}
		fmt.Println("Неверное значение! Введите номер месяца от 1 до 12: ")
	}

	fmt.Println("---------------------")

	switch month {
	case 1, 2, 12:
			fmt.Println("Зима")

	case 3, 4, 5:
		fmt.Println("Весна")

	case 6, 7, 8:
		fmt.Println("Лето")

	case 9, 10, 11:
		fmt.Println("Осень")

	default:
		fmt.Println("Что-то другое")
	}

	fmt.Println("---------------------")
}
