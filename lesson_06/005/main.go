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
	firstBasketOpen := true
	secondBasketOpen := true

	// for {
	// 	check := firstBasket < 6
	// 	if check {
	// 		firstBasket++
	// 	}

	// 	check = secondBasket < 4
	// 	if check {
	// 		secondBasket++
	// 	}

	// 	check = thirdBasket < 9
	// 	if check {
	// 		thirdBasket++
	// 		continue
	// 	}

	// 	break
	// }

	for {
		if firstBasketOpen && firstBasket < 6 {
			firstBasket++
			firstBasketOpen = false
			continue
		}

		if secondBasketOpen && secondBasket < 4 {
			secondBasket++
			secondBasketOpen = false
			continue
		}

		if thirdBasket < 9 {
			thirdBasket++
			firstBasketOpen = true
			secondBasketOpen = true
			continue
		}

		break
	}

	fmt.Println("В первой корзине:", firstBasket)
	fmt.Println("Во второй корзине:", secondBasket)
	fmt.Println("В тертьей корзине:", thirdBasket)
}
