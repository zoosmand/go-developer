package main

import (
	"fmt"
	"math"

	usrcons "../../packages/usrcons"
)

func main() {

	h := "Задача 10.1. Разложение ex в ряд Тейлора."
	usrcons.PrintHeader(&h)

	// ----------------------------------------------------
	var consoleInput int
	invite := "Введите степень и количество знаков после запятой."
	fmt.Println(invite)

	var d []int
	var c int

	// ----------------------------------------------------
	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			d = append(d, consoleInput)
			c++
		}
		if c == 2 {
			break
		}
	}

	fmt.Println("-----------------------------")
	fmt.Print("e: ")
	fmt.Println(CalcE(&d))
	fmt.Println("-----------------------------")
}

func CalcE(d *[]int) float64 {
	x := (*d)[0]
	n := (*d)[1]	 
	precision := 1 / math.Pow(float64(10), float64(n))

	var myE float64
	var preE float64

	fact := 1
	myE = 1
	i := 1

	for {
		fact *= i
		myE += math.Pow(float64(x), float64(i)) / float64(fact)
		i++
		if myE-preE < precision {
			break
		}
		preE = myE
	}
	return myE
}
