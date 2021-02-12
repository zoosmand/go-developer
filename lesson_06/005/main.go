package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.5. Наполнение корзин\n\n")

	var firstBasket int
	var secondBasket int
	var thirdBasket int

	for {
		check := firstBasket < 6
		if check {
			firstBasket++
		}

		check = secondBasket < 4
		if check {
			secondBasket++
		}

		check = thirdBasket < 9
		if check {
			thirdBasket++
			continue
		}

		break
	}

	fmt.Println("В первой корзине:", firstBasket)
	fmt.Println("Во второй корзине:", secondBasket)
	fmt.Println("В тертьей корзине:", thirdBasket)
}
