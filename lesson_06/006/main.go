package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print("Задача 6.6. Движение лифта\n\n")

	liftCapacity := 2

	tenthDeckPassengers := 3
	seventhDeckPassengers := 3
	forthDeckPassengers := 3

	for {
		allPassengers := tenthDeckPassengers + seventhDeckPassengers + forthDeckPassengers
		if allPassengers <= 0 {
			break
		}

		fmt.Println("Лифт начал движение вниз с 24-го этажа")
		passengersInLift := 0
		for i := 24; i > 0; i-- {

			if passengersInLift == liftCapacity {
				continue
			}

			if i == 10 {
				for tenthDeckPassengers != 0 {
					if passengersInLift == liftCapacity {
						break
					}
					passengersInLift++
					tenthDeckPassengers--
					fmt.Println("Берем пассажира с десятого этажа")
				}
			}

			if i == 7 {
				for seventhDeckPassengers != 0 {
					if passengersInLift == liftCapacity {
						break
					}
					passengersInLift++
					seventhDeckPassengers--
					fmt.Println("Берем пассажира с седьмого этажа")
				}
			}

			if i == 4 {
				for forthDeckPassengers != 0 {
					if passengersInLift == liftCapacity {
						break
					}
					passengersInLift++
					forthDeckPassengers--
					fmt.Println("Берем пассажира с четвертого этажа")
				}
			}
		}
		fmt.Print("Лифт доставил пассажиров на первый этаж и поднимается наверх\n\n")
	}

	fmt.Println("Лифт закончил работу")
}
