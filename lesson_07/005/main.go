package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.5. Угадай число\n\n")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(101)

	var number int
	fmt.Println("Введите номер от 1 до 100: ")

	for {
		_, _ = fmt.Scan(&number)

		if number >= 1 && number <= 100 {
			if secretNumber > number {
				fmt.Println("Больше!")
			} else if secretNumber < number {
				fmt.Println("Меньше!")
			} else if secretNumber == number {
				fmt.Println("Найдено!")
				break
			}

		} else {
			fmt.Println("Неверное значение! Введите номер от 1 до 100: ")
		}
	}

}
