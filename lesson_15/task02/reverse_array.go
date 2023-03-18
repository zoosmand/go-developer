package task02

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func ReverseIntArray() {
	h := "Задача 15.2. Функция, реверсирующая массив."
	fmt.Print(usecons.Header(&h))

	const arrLen int = 10
	var digits [arrLen]int

	// ----------------------------------------------------
	var consoleInput int
	invite := fmt.Sprintf("Введите %v чисел. (последовательно или одной строкой через запятую)", arrLen)
	fmt.Println(invite)

	// ----------------------------------------------------
	i := 0
	for {
		_, err := fmt.Scan(&consoleInput)
		if err != nil {
			fmt.Println("Неверный ввод!", invite)
		} else {
			digits[i] = consoleInput
			i++
			if i >= arrLen {
				break
			}
		}
	}

	fmt.Printf("\nИсходный массив: %v\n", digits)
	reverseIntArray(digits[0:])
	fmt.Printf("Результат:       %v\n\n", digits)
}

func reverseIntArray(arr []int) {
	l := len(arr)

	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}
