package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.7. Счастливый билет")
	fmt.Println()

	fmt.Println("Введите номер билета: ")
	var bilette int
	fmt.Scan(&bilette)
	fmt.Println()

	segment0 := bilette/1000
	segment1 := bilette%1000/100
	segment2 := bilette%100/10
	segment3 := bilette%10

	message := ""

	if (segment0 + segment1) == (segment2 + segment3) {
		if (segment0 == segment3) && (segment1 == segment2) {
			message = "Это зеркальный счастливый билет."
		} else {
			message = "Это счастливый билет."
		}
	} else {
		message = "Это обычный билет."
	}

	fmt.Println(message)
	fmt.Println()
}
