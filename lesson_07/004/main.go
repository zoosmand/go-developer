package main

import (
	"fmt"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 7.4. Максимальное расстояние между двумя счастливыми билетами\n\n")

	distanceNextHappyTicket := 0
	growingDistance := 1

	for i := 100; i <= 999; i++ {
		// take digits from the iteration number "i", then sum them
		m := i/100 + i%100/10 + i%10

		for j := 0; j <= 999; j++ {

			// take digits from the iteration number "j", then sum them
			k := j/100 + j%100/10 + j%10

			// increase distance counter if sums compare
			// otherwise means the distance has found
			if m != k {
				distanceNextHappyTicket++
			} else {
				if growingDistance < distanceNextHappyTicket {

					// increase the found number to 1 to get the right arithmetic number
					growingDistance = distanceNextHappyTicket + 1
					fmt.Println(growingDistance)
				}
				distanceNextHappyTicket = 0
			}
		}
	}
}
