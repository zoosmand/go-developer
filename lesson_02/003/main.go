package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите длительность смены в минутах: ")
	var workingDayMin int = 480
	fmt.Println(workingDayMin)

	fmt.Println("Сколько минут клиент делает заказ?")
	var timeToOrder int = 2
	fmt.Println(timeToOrder)

	fmt.Println("Сколько минут кассир собирает заказ?")
	var timeToCombine int = 4
	fmt.Println(timeToCombine)

	fmt.Println("-----Считаем-----")
	fmt.Println(
		"За смену длиной",
		workingDayMin,
		"минут кассир успеет обслужить",
		workingDayMin/(timeToCombine+timeToOrder),
		"клиентов.",
	)
}
