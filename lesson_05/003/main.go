package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.3. Проверка на совпадение")
	fmt.Println()

	var isEqual bool = false

	var firstDigit int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&firstDigit)

	var secondDigit int
	fmt.Println("Введите второе число: ")
	fmt.Scan(&secondDigit)

	var thirdDigit int
	fmt.Println("Введите третье число: ")
	fmt.Scan(&thirdDigit)

	if (firstDigit == secondDigit) || (firstDigit == thirdDigit) || (secondDigit == thirdDigit) {
		isEqual = true
	}


	if isEqual {
		fmt.Println("Совпадение найдено.")
	} else {
		fmt.Println("Совпадений нет.")
	}
	fmt.Println()
}
