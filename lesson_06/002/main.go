package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.2. Сумма двх чисел по единице\n\n")

	var firstNumber int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&firstNumber)
	fmt.Println()

 	var secondNumber int
	fmt.Println("Введите второе число: ")
	fmt.Scan(&secondNumber)
	fmt.Println()

	for i := firstNumber; i < (firstNumber+secondNumber); i++ {
		fmt.Println(i)
	}

	fmt.Println()
}
