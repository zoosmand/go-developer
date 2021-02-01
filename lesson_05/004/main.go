package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.4. Сумма без сдачи")
	fmt.Println()

	nominals := []int{1, 2, 5} 		// available nominals
	var firstCoin int
	var secondCoin int
	var thirdCoin int
	var isConformed bool = false 	// check conformation amount and coins
	var coinIsOK bool = false 		// check coin's nominal during input 

	// get the amount
	var amount int
	fmt.Println("Введите сумму: ")
	fmt.Scan(&amount)

	// chech coins until they have put right
	for !coinIsOK {
		fmt.Println("Введите номинал первой монеты: ")
		fmt.Scan(&firstCoin)
		for _, v := range nominals {
			if firstCoin == v {
				coinIsOK = true
				break
			} else {
				coinIsOK = false
			}
		}

		fmt.Println("Введите номинал второй монеты: ")
		fmt.Scan(&secondCoin)
		for _, v := range nominals {
			if secondCoin == v {
				coinIsOK = true
				break
			} else {
				coinIsOK = false
			}
		}

		fmt.Println("Введите номинал третьей монеты: ")
		fmt.Scan(&thirdCoin)
		for _, v := range nominals {
			if thirdCoin == v {
				coinIsOK = true
				break
			} else {
				coinIsOK = false
			}
		}

		if !coinIsOK {
			fmt.Println("Номиналы монет не верны, введите номималы еще раз: ")
			fmt.Println()
		}
	}

	fmt.Println()

	// chech if one coin could be used
	if (amount == firstCoin) || (amount == secondCoin) || (amount == thirdCoin) {
		// OK
		isConformed = true
	}

	// chech if two coins could be used
	if (amount == (firstCoin + secondCoin)) || (amount == (firstCoin + thirdCoin)) || (amount == (secondCoin + thirdCoin)) {
		//  OK
		isConformed = true
	}

	// chech if all coins could be used
	if (amount == (firstCoin + secondCoin + thirdCoin)) {
		//  OK
		isConformed = true
	}

	if isConformed {
		fmt.Println("Пользователь может оплатить без сдачи.")
		fmt.Println()
	} else {
		fmt.Println("Пользователь не может оплатить без сдачи.")
		fmt.Println()
	}
}
