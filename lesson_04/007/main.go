package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.7. Прогрессивный налог")
	fmt.Println()

	var totalIncome int
	fmt.Println("Введите сумму вашего совокупного дохода:")
	fmt.Scan(&totalIncome)

	var totalTax int

	if totalIncome > 50000 {
		// totalTax = (totalIncome-50000)/100*30 + 40000/100*20 + 10000/100*13
		totalTax = (totalIncome-50000)/100*30 + 9300
	} else if totalIncome > 10000 && totalIncome <= 50000 {
		// totalTax = (totalIncome-10000)/100*20 + 10000/100*13
		totalTax = (totalIncome-10000)/100*20 + 1300
	} else if totalIncome <= 10000 {
		totalTax = totalIncome/100*13
	}

	fmt.Println()
	fmt.Print("Сумма налога к оплате ", totalTax, " рублей.\n")
	fmt.Println()
}
