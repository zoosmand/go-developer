package task03

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func Returns() {
	h := "Задача 14.3. Именованные возвращаемые значения."
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

	fmt.Printf("Результат умножения на 100: %d\n", multiplicationBy100(consoleInput))
	fmt.Printf("Результат сложения с числом 100: %d\n\n", additionOn100(consoleInput))
}

func multiplicationBy100(dig int) (result int) {
	result = dig * 100
	return result
}

func additionOn100(dig int) (result int) {
	result = dig + 100
	return result
}
