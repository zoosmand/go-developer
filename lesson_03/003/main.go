package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 3. Обмен местами.")
	fmt.Println()

	a := 42
	b := 153

	fmt.Println("До:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()

	tmp := a
	a = b
	b = tmp

	fmt.Println("После:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println()
}
