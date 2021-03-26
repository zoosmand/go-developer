package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.5. Угадай число\n\n")

	var number int
	fmt.Println("Введите номер от 1 до 100: ")

	for {
		_, _ = fmt.Scan(&number)

		if number >= 1 && number <= 100 {
			break
		}
		fmt.Println("Неверное значение! Введите номер от 1 до 100: ")
	}

}
