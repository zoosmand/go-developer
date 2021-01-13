package main

import (
	"fmt"
)

const lowThreshold int = 100
const highThreshold int = 100000

func main() {
	fmt.Println()
	fmt.Println("Задача 4.3. Банкомат")
	fmt.Println()

	var withdrawSum int
	fmt.Println("Введите сумму для снятия не меньше", lowThreshold, "и не больше", highThreshold, "рублей:")
	fmt.Scan(&withdrawSum)

	fmt.Println()

	if withdrawSum < lowThreshold {
		fmt.Print("Запрашиваемая сумма меньше минимального порога в ", lowThreshold, " рублей.\n")
	} else if withdrawSum > highThreshold {
		fmt.Print("Запрашиваемая сумма превышает порог в ", highThreshold, " рублей.\n")
	} else {
		if (withdrawSum % 100) != 0 {
			fmt.Print(
				"По првилам нашего банка сумма должна быть кратна 100 рублям,\n",
				"поэтому мы уменьшили суммы до кратного значения в ", withdrawSum/100*100, " рублей.\n",
			)
		}
		fmt.Println("Возьмите ваши деньги. Спасибо за пользование услугами нашего банка.")
	}

	fmt.Println()
}
