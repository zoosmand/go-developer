package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.3 Злостные вредители.")
	fmt.Println()

	var growingSpeed int
	fmt.Println("Введите скорость роста бамбука в см/день:")
	fmt.Scan(&growingSpeed)

	var wormEatingSpeed int
	fmt.Println("Введите скорость поедания гусеницами в см/день:")
	fmt.Scan(&wormEatingSpeed)

	var sproutHeight int
	fmt.Println("Введите высоту черенка в см:")
	fmt.Scan(&sproutHeight)

	var numberOfDays int
	fmt.Println("Введите количество дней:")
	fmt.Scan(&numberOfDays)

	totalHeight := sproutHeight + (growingSpeed-wormEatingSpeed)*numberOfDays

	fmt.Print("Через ", numberOfDays, " день(дня, дней) высота бамбука составит ", totalHeight, "см.\n")
	fmt.Println()

}
