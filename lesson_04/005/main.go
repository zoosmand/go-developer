package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.5. Ресторан")
	fmt.Println()

	var dayWeek int
	fmt.Print("Введите день недели:\n",
		"\t1 - Понедельник\n",
		"\t2 - Вторник\n",
		"\t3 - Среда\n",
		"\t4 - Четверг\n",
		"\t5 - Пятница\n",
		"\t6 - Суббота\n",
		"\t7 - Воскресенье\n")
	fmt.Scan(&dayWeek)

	var guestCount int
	fmt.Println("Введите количество гостей:")
	fmt.Scan(&guestCount)

	var checkSum int
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checkSum)

	// // check the Monday discount
	// if dayWeek == 1 {
	// 	checkSum -= int(float32(checkSum) * 0.1)
	// }

	// // check the Friday discount for the more than 10000
	// if dayWeek == 5 && checkSum >= 10000 {
	// 	checkSum -= int(float32(checkSum) * 0.05)
	// }

	// // check guest count on any day
	// if guestCount > 5 {
	// 	checkSum += int(float32(checkSum) * 0.1)
	// }

	// fmt.Println()
	// fmt.Print("Сумма к оплате ", checkSum, " рублей.\n")
	// fmt.Println()

	// Следуя требованиям задачи, в понедельник компания из 5 человек
	// получит скидку в 10% и надбавку за обслуживание в 10%
	// это может привести к искажению результата суммы к оплате
	// 
	// Суммируя скидки перед вычислением сумммы к оплате,
	// можно избавится от искажений

	var totalDiscount int

	// check the Monday discount
	if dayWeek == 1 {
		totalDiscount -= 10
	}

	// check the Friday discount for the more than 10000
	if dayWeek == 5 && checkSum >= 10000 {
		totalDiscount -= 5
	}

	// check guest count on any day
	if guestCount > 5 {
		totalDiscount += 10
	}

	fmt.Println()
	fmt.Print("Сумма к оплате ", checkSum + int(checkSum * totalDiscount / 100), " рублей.\n")
	fmt.Println()
}
