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

	switch month {
	case 1:
	case 2:
	case 12:
		fmt.Println("Зима")

	case 3:
	case 4:
	case 5:
		fmt.Println("Весна")

	case 6:
	case 7:
	case 8:
		fmt.Println("Лето")

	case 9:
	case 10:
	case 11:
		fmt.Println("Осень")

	default:
		fmt.Println("Что-то другое")
	}

	fmt.Println("---------------------")
}
