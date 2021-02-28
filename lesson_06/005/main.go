package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.5. Наполнение корзин\n\n")

	firstBasket := 0
	secondBasket := 0
	thirdBasket := 0

	for {
		if firstBasket < 6 {
			firstBasket++
			continue
		}

		if secondBasket < 4 {
			secondBasket++
			continue
		}

		if thirdBasket < 9 {
			thirdBasket++
			continue
		}

		break
	}

	fmt.Println("В первой корзине:", firstBasket)
	fmt.Println("Во второй корзине:", secondBasket)
	fmt.Println("В тертьей корзине:", thirdBasket)
}
