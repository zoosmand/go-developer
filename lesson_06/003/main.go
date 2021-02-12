package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.1. Расчет суммы скидки\n\n")

	areRight := false

	var price int
	var discount int
	var finalDiscount int

	fmt.Println("Введите цену товара: ")
	fmt.Scan(&price)

	for !areRight {
		if discount != 0 {
			fmt.Print("Скидка слишком велика, повторите ввод.\n\n")
		}

		fmt.Println("Введите сумму скидки: ")
		fmt.Scan(&discount)

		finalDiscount = (price / 100) * discount

		if finalDiscount > 2000 {
			finalDiscount = 2000
		}

		if discount <= 30 {
			areRight = true
		}
	}

	fmt.Print("\nСумма скидки: ", finalDiscount, "\n")
}
