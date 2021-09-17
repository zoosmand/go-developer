package main

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func main() {
	h := "Задача 12.7. Скобки"
	fmt.Print(usecons.Header(&h))

	// ----------------------------------------------------
	invite := "Введите число."
	fmt.Println(invite)

	// ----------------------------------------------------
	var num int
	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			break
		}
	}

	fmt.Println("------------------------------------------------")

	simple()
	fmt.Println("------------------------------------------------")

	embrace(num)
}

func embrace(num int) {

	var eLeft []byte
	var eRight []byte
	for i := 0; i < num; i++ {
		eLeft = append(eLeft, '(')
		eRight = append(eRight, '(')
	}
	for i := 0; i < num; i++ {
		eLeft = append(eLeft, ')')
		eRight = append(eRight, ')')
	}

	counter := 1
	fmt.Printf("% 4.d - %s\r\n", counter, eLeft)

	for k := 1; k < num; k++ {
		for i := -k; i < num-k*2; i++ {
			eLeft[num+i], eLeft[num+i+1] = eLeft[num+i+1], eLeft[num+i]
			counter++
			fmt.Printf("% 4.d - %s\r\n", counter, eLeft)
			eRight[num-i-1], eRight[num-i-2] = eRight[num-i-2], eRight[num-i-1]
			if string(eLeft) != string(eRight) {
				counter++
				fmt.Printf("% 4.d - %s\r\n", counter, eRight)
			}
		}
	}
}
