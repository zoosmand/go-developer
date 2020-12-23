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

	var targetHeight int
	fmt.Println("Введите целевую высоту бамбука:")
	fmt.Scan(&targetHeight)

	totalDays := (targetHeight - sproutHeight) / (growingSpeed - wormEatingSpeed)

	fmt.Print("Для достижения целевой высоты в ", targetHeight, "см поанадобится ", totalDays, " день(дня, дней)\n")
	fmt.Println()
}
