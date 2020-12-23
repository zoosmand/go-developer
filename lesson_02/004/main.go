package main

import (
	"fmt"
)

func main() {
	var totalCommunityExpenses int = 4000000
	fmt.Println("Сумма, указанная в квитанции:", totalCommunityExpenses)

	var numberOfPorche int = 10
	fmt.Println("Подъездов в доме:", numberOfPorche)

	var numberFlatInOnePorch int = 4
	fmt.Println("Квартир в каждом подъезде:", numberFlatInOnePorch)

	fmt.Println("----Результат-----")
	fmt.Println(
		"Каждая квартира должна заплатить по",
		totalCommunityExpenses/(numberFlatInOnePorch*numberOfPorche),
		"руб.",
	)
}
