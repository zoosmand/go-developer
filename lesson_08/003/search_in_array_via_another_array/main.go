package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 8.3.extra Array iretations \n\n")

	source := [3]int{5, 10, 20}

	checkValues := [6]int{5, 7, 20, 77, 6, 10}

	for i := 0; i < len(checkValues); i++ {
		checkFlag := 0
		for j := 0; j < len(source); j++ {
			if checkValues[i] == source[j] {
				checkFlag |= 1
				break
			}
		}
		fmt.Println(checkFlag)
	}

	fmt.Println("---------------------")

	fmt.Println("---------------------")
}
