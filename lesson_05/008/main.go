package main

import (
	"fmt"
)

func main() {
	// Introdunction
	fmt.Println()
	fmt.Print("Задача 5.8. Угадывание числа\n\n")

	fmt.Print("Загадайте любое число от 1 до 10.\n\n")
	fmt.Print("Введите '0' если число угадано, '>' - если число больше загаданного или '<' если оно меньше загаданного: \n\n")

	// initiate vars
	secretDigit := 5
	var inputIsRight bool = false
	var userInput rune = 48

	// define user allowed input symbols
	allowedSymbols := [3]rune{'0', '<', '>'}

	// make iteractions
	for i := 0; i < 4; i++ {
		fmt.Printf("Вы загадали число %d?\n", secretDigit)

		// wait until user input gets to allow
		for !inputIsRight {
			fmt.Println("---")
			fmt.Scanf("%c", &userInput)
			for _, v := range allowedSymbols {
				if v == userInput {
					inputIsRight = true
					break
				} else {
					inputIsRight = false
				}
			}
		}

		// is the digit is less than the one that was diaplayed, decrement it by one
		if userInput == '<' {
			secretDigit--
		// is the digit is more than the one that was diaplayed, increment it by two
		} else if userInput == '>' {
			secretDigit += 2
			// check overload
			if secretDigit > 10 {
				secretDigit = 10
			}
		// if OK, out the program
		} else if userInput == '0' {
			break
		}

		// clear the flag, that checks user input
		inputIsRight = false
	}

	fmt.Printf("\nУгаданное число: %d.\n\n", secretDigit)
}
