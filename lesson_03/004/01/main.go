package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.1 Злостные вредители.")
	fmt.Println()

	growingSpeed := 50
	wormEatingSpeed := 20
	sproutHeight := 100

	totalHeight := sproutHeight + (growingSpeed-wormEatingSpeed)*2 + growingSpeed/2

	fmt.Print("Через 3.5 дня высота бамбука составит ", totalHeight, "см.\n")
	fmt.Println()

}
