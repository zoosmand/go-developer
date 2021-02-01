package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.2. Проверка на положительное значение числа")
	fmt.Println()

	var isPositive bool = false

	var firstDigit int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&firstDigit)

	var secondDigit int
	fmt.Println("Введите второе число: ")
	fmt.Scan(&secondDigit)

	var thirdDigit int
	fmt.Println("Введите третье число: ")
	fmt.Scan(&thirdDigit)

	if (firstDigit > 0) || (secondDigit > 0) || (thirdDigit > 0) {
		isPositive = true
	}


	if isPositive {
		fmt.Println("Как минимум одно число является положительным.")
		fmt.Println()
	} else {
		fmt.Println("Ни одно число из введенных не является положительным.")
		fmt.Println()
	}
}
