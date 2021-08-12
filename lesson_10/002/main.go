package main

import (
	"fmt"
	"math"

	usrcons "../../packages/usrcons"
)

func main() {
	h := "Задача 10.2. Рассчет размена банковского депозита."
	usrcons.PrintHeader(&h)

	// ----------------------------------------------------
	var consoleInput int
	invite := "Введите сумму вклада, процент капитазации, срок (в годах)."
	fmt.Println(invite)

	var d []int
	var c int

	// ----------------------------------------------------
	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			d = append(d, consoleInput)
			c++
		}
		if c == 3 {
			break
		}
	}
	calculateInterest(&d)
}

func calculateInterest(d *[]int) {
	months := 12 * (*d)[2]
	deposit := float64((*d)[0])
	interest := float64((*d)[1])

	bankInterest := 0.0

	for i := 0; i < months; i++ {
		clearInterest := deposit * (interest / 100)
		clientInterest := math.Floor(clearInterest*100) / 100
		bankInterest += clearInterest - clientInterest

		deposit += clientInterest
	}

	fmt.Println("-----------------------------")
	fmt.Print("Размер депозита: ")
	fmt.Print(math.Round(deposit*100) / 100)
	fmt.Print("\nДоход банка : ")
	fmt.Println(bankInterest)
	fmt.Println("-----------------------------")
}
