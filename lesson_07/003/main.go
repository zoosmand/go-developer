package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.3. Елочка\n\n")

	var height int
	fmt.Println("Введите высоту елочки: ")
	fmt.Scan(&height)

	fmt.Println("---------------------")

	for i := 0; i < height; i++ {
		var branch []rune
		// left branch
		for k := 0; k < height+1-i; k++ {
			branch = append(branch, ' ')
		}
		for k := 0; k < i+1; k++ {
			branch = append(branch, '*')
		}
		// right branch
		for k := 0; k < i; k++ {
			branch = append(branch, '*')
		}
		for k := 0; k < height-i; k++ {
			branch = append(branch, ' ')
		}
		fmt.Println(string(branch))
	}

	fmt.Println("---------------------")
}
