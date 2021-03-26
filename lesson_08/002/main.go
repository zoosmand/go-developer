package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 8.2.\n\n")

	workingDaysShort := [5]string{"пн", "вт", "ср", "чт", "пт"}

	var inputDayName string
	fmt.Println("Введите день", workingDaysShort)

checkInputLoop:
	for {
		_, _ = fmt.Scan(&inputDayName)

		for i := 0; i < len(workingDaysShort); i++ {
			if workingDaysShort[i] == inputDayName {
				break checkInputLoop
			}
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
