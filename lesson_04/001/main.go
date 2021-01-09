package main

import (
	"fmt"
)

const threshholdAccept int = 275

func main() {
	fmt.Println()
	fmt.Println("Задача 4.1. Баллы ЕГЭ")
	fmt.Println()

	var mathMark int
	fmt.Println("Введите количество баллов по математике:")
	fmt.Scan(&mathMark)

	var physicsMark int
	fmt.Println("Введите количество баллов по физике:")
	fmt.Scan(&physicsMark)

	var chemistryMark int
	fmt.Println("Введите количество баллов по химии:")
	fmt.Scan(&chemistryMark)

	totalScore := mathMark + physicsMark + chemistryMark

	if totalScore >= threshholdAccept {
		fmt.Print("Сумма ваших баллов ", totalScore, ". Поздравляю, вы зачислены.\n")
	} else {
		fmt.Print("Сожалею, но ам не хватает для зачисления всего ", threshholdAccept-totalScore, " балла(балл, баллов). Удачи на следующий раз.\n")
	}

	fmt.Println()

}
