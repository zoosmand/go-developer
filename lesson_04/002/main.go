package main

import (
	"fmt"
)

const threshold int = 5

func main() {
	fmt.Println()
	fmt.Println("Задача 4.2. Три числа")
	fmt.Println()

	var firstDigit int
	fmt.Println("Введите первое число:")
	fmt.Scan(&firstDigit)

	var secondDigit int
	fmt.Println("Введите второе число:")
	fmt.Scan(&secondDigit)

	var firdDigit int
	fmt.Println("Введите третье число:")
	fmt.Scan(&firdDigit)

	fmt.Println()

	if firstDigit > threshold {
		fmt.Print("Первое число ", firstDigit, " больше чем ", threshold, ".\n")
	} else if secondDigit > threshold {
		fmt.Print("Второе число ", secondDigit, " больше чем ", threshold, ".\n")
	} else if firdDigit > threshold {
		fmt.Print("Третье число ", firdDigit, " больше чем ", threshold, ".\n")
	} else {
		fmt.Print("Все введенные числа меньше чем ", threshold, ".\n")
	}

	fmt.Println()
}
