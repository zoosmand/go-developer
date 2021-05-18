package main

import (
	"fmt"
	"math"
)

func main() {
	// Introduction
	fmt.Print("\nЗадача 9.1. Переполнение.\n")
	fmt.Print("-------------------------\n\n")

	var counterUint8 uint8
	var counterUint16 uint16

	var counter8 uint32
	var counter16 uint32

	for i := 0; i < math.MaxUint32; i++ {
		counterUint8++
		counterUint16++

		if counterUint8 == 0 {
			counter8++
		}
		if counterUint16 == 0 {
			counter16++
		}
	}

	fmt.Println("Количество переполнений 8-ми битного беззнакового значения в 32-х битном беззнаковом поле: ", counter8+1)
	fmt.Println("Количество переполнений 16-ти битного беззнакового значения в 32-х битном беззнаковом поле: ", counter16+1)
}
