package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.2 Злостные вредители.")
	fmt.Println()

	growingSpeed := 50
	wormEatingSpeed := 20
	sproutHeight := 100
	targetHeight := 300

	totalDays := (targetHeight - sproutHeight) / (growingSpeed - wormEatingSpeed)

	fmt.Print("Для достижения целевой высоты в ", targetHeight, "см поанадобится ", totalDays, " день(дня, дней)\n")
	fmt.Println()

}
