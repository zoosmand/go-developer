package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.4. Максимальное расстояние между двумя счастливыми билетами\n\n")

	distanceNextHappyTicket := 0
	prevHappyTicket := 100000

	for i := 100000; i <= 999999; i++ {
		head := i / 1000
		tail := i % 1000
		headSum := head/100 + head%10 + head/10%10
		tailSum := tail/100 + tail%10 + tail/10%10

		if headSum == tailSum {
			if distanceNextHappyTicket < i-prevHappyTicket {
				distanceNextHappyTicket = i - prevHappyTicket
				// fmt.Println(distanceNextHappyTicket)
			}
			prevHappyTicket = i
		}

	}

	fmt.Println(distanceNextHappyTicket)
}
