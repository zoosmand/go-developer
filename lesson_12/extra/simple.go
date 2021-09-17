package main

import (
	"fmt"
)

func simple() {

	num := 4

	var eLeft []byte
	var eRight []byte
	for i := 0; i < num; i++ {
		eLeft = append(eLeft, '1')
		eRight = append(eRight, '1')
	}
	for i := 0; i < num; i++ {
		eLeft = append(eLeft, '0')
		eRight = append(eRight, '0')
	}

	// counter := 1
	// fmt.Printf("% 4.d - %s\r\n", counter, eLeft)

	// for k := 1; k < num; k++ {
	// 	for i := -k; i < num-k*2; i++ {
	// 		eLeft[num+i], eLeft[num+i+1] = eLeft[num+i+1], eLeft[num+i]
	// 		counter++
	// 		fmt.Printf("% 4.d - %s\r\n", counter, eLeft)
	// 		eRight[num-i-1], eRight[num-i-2] = eRight[num-i-2], eRight[num-i-1]
	// 		if string(eLeft) != string(eRight) {
	// 			counter++
	// 			fmt.Printf("% 4.d - %s\r\n", counter, eRight)
	// 		}
	// 	}
	// }

	// Количество итераций n-1
	fmt.Printf("%s\r\n", eLeft)

	// первая итерация в n-1 проходов: меняем местами краюнюю правую единицу и ближайщий ноль
	eLeft[num-1], eLeft[num] = eLeft[num], eLeft[num-1]
	fmt.Printf("%s\r\n", eLeft) // 11101000

	eLeft[num], eLeft[num+1] = eLeft[num+1], eLeft[num]
	fmt.Printf("%s\r\n", eLeft) // 11100100

	eLeft[num+1], eLeft[num+2] = eLeft[num+2], eLeft[num+1]
	fmt.Printf("%s\r\n", eLeft) // 11100010

	// вторая итерация итерация в n-2 проходов: меняем местами вторую справа единицу и ближайший ноль,
	eLeft[num-2], eLeft[num-1] = eLeft[num-1], eLeft[num-2]
	fmt.Printf("%s\r\n", eLeft) // 11010010

	eLeft[num-1], eLeft[num] = eLeft[num], eLeft[num-1]
	fmt.Printf("%s\r\n", eLeft) // 11001010

	// третья итерация итерация в n-3 проходов: меняем местами третью справа единицу и ближайший ноль,
	eLeft[num-3], eLeft[num-2] = eLeft[num-2], eLeft[num-3]
	fmt.Printf("%s\r\n", eLeft) // 10101010

	// Одновременно с проходом слева-направо, делайм проход справа-налево. Если массивы не равны, то выводим на печать.

}
