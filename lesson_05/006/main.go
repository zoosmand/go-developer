package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 5.6. Решение квадратного уравнения")
	fmt.Println()

	var coeffs []float64

	fmt.Println("Введите последовательно коэффициенты a, b и c: ")
	for i := 0; i < 3; i++ {
		fmt.Println("---")
		var coeff float64
		fmt.Scan(&coeff)
		coeffs = append(coeffs, coeff)
	}
	fmt.Println()

	discriminant := math.Pow(coeffs[1], 2) - 4*coeffs[0]*coeffs[2]
	message := ""

	if discriminant < 0 {
		message = "Действительных корней нет."
	} else if discriminant == 0 {
		answer := -1 * coeffs[1] / 2 * coeffs[0]
		message = "Один корень: " + fmt.Sprintf("%.2f", answer)
	} else {
		answer1 := (-1*coeffs[1] + math.Sqrt(discriminant)) / 2 * coeffs[0]
		answer2 := (-1*coeffs[1] - math.Sqrt(discriminant)) / 2 * coeffs[0]
		message = "Два корня: " + fmt.Sprintf("%.2f", answer1) + " и " + fmt.Sprintf("%.2f", answer2)
	}

	fmt.Println(message)
	fmt.Println()
}
