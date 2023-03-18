package task04

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

var globalVar1 = 1234
var globalVar2 = 5678
var globalVar3 = 9012

func Returns() {
	h := "Задача 14.4. Именованные возвращаемые значения."
	fmt.Print(usecons.Header(&h))

	// ----------------------------------------------------
	var consoleInput int
	invite := "Введите целое число."
	fmt.Println(invite)

	// ----------------------------------------------------
	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			break
		}
	}

	fmt.Printf("Результат func1: %d\n", func1(consoleInput))
	fmt.Printf("Результат func2: %d\n", func2(consoleInput))
	fmt.Printf("Результат func3: %d\n", func3(consoleInput))

	fmt.Printf("Результат последовательного вызова: %d\n\n", func1(func2(func3(consoleInput))))

}

func func1(dig int) int {
	return dig + globalVar1
}

func func2(dig int) int {
	return dig + globalVar2
}

func func3(dig int) int {
	return dig + globalVar3
}
