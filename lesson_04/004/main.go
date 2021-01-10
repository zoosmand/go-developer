package main

import (
	"fmt"
)

const threshold int = 5

func main() {
	fmt.Println()
	fmt.Println("Задача 4.4. Три числа. Еще попытка")
	fmt.Println()

	var digitCount int
	var inputDigit int

	// Здесь напрашивается цикл и массив с именами чисел(цифр)
	fmt.Println("Введите первое число:")
	fmt.Scan(&inputDigit)
	
	if inputDigit >= threshold {
		digitCount++
	}

	fmt.Println("Введите второе число:")
	fmt.Scan(&inputDigit)

	if inputDigit >= threshold {
		digitCount++
	}

	fmt.Println("Введите третье число:")
	fmt.Scan(&inputDigit)

	if inputDigit >= threshold {
		digitCount++
	}

	fmt.Println()
	fmt.Print("Количество введенных чисел больше ", threshold, " - ", digitCount, ".\n")
	fmt.Println()
}
