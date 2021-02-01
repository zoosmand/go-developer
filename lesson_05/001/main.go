package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.1. Определение координатной плоскости точки")
	fmt.Println()

	var abscissa int
	fmt.Println("Введите координату оси абсцисс:")
	fmt.Scan(&abscissa)

	var ordinate int
	fmt.Println("Введите координату оси ординат:")
	fmt.Scan(&ordinate)

	if (abscissa > 0) && (ordinate > 0) {
		fmt.Println("Точка находится в первой четверти.")
		fmt.Println()
	} else if (abscissa < 0) && (ordinate > 0) {
		fmt.Println("Точка находится во второй четверти.")
		fmt.Println()
	} else if (abscissa < 0) && (ordinate < 0) {
		fmt.Println("Точка находится в третьей четверти.")
		fmt.Println()
	} else if (abscissa > 0) && (ordinate < 0) {
		fmt.Println("Точка находится в четвертой четверти.")
		fmt.Println()
	}
}
