package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 8.2.\n\n")

	workingDaysShort := map[string]string{
		"пн": "понедельник",
		"вт": "вторник",
		"ср": "средв",
		"чт": "четверг",
		"пт": "пятница",
	}

	var inputDayName string
	// fmt.Println("Введите день [ пн | вт | ср | чт | пт ]: ")
	fmt.Println("Введите день", workingDaysShort)

checkInputLoop:
	for {
		_, _ = fmt.Scan(&inputDayName)
		if _, ok := workingDaysShort[inputDayName]; ok {
			break checkInputLoop
		}

		fmt.Println("Неверное значение! Введите день", workingDaysShort)
	}

	fmt.Println("---------------------")

	switch inputDayName {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	}

	fmt.Println("---------------------")
}
