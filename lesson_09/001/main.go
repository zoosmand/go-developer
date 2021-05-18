package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// Introduction
	fmt.Print("\nЗадача 9.1. Переполнение.\n")
	fmt.Print("-------------------------\n\n")

	counter8 := uint32(0xffffffff) >> uint32(bits.Len32(uint32(0xff)))
	counter16 := uint32(0xffffffff) >> uint32(bits.Len32(uint32(0xffff)))

	fmt.Println("Количество переполнений 8-ми битного беззнакового значения в 32-х битном беззнаковом поле: ", counter8)
	fmt.Println("Количество переполнений 16-ти битного беззнакового значения в 32-х битном беззнаковом поле: ", counter16)
}
