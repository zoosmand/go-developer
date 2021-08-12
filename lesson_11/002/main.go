package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	usrcons "../../packages/usrcons"
)

func main() {

	h := "Задача 11.2. Строки."
	usrcons.PrintHeader(&h)

	myLine := "a10 10 20b 20 30c30 30 dd"
	// myLine := "     "
	// myLine := ""
	myArray := strings.Split(myLine, " ")
	myArraySize := len(myArray)
	rule := regexp.MustCompile(`[a-f]`)

	fmt.Println("------------------------------------------------")
	for i := 0; i < myArraySize; i++ {
		base := 10
		if rule.MatchString(myArray[i]) {
			base = 16
		}

		if myArray[i] == "" {
			fmt.Println("Невозможно преобразовать пустое значение")
		} else {
			a, err := strconv.ParseInt(myArray[i], base, 0)
			if err != nil {
				fmt.Println("Не удалось преобразовать строку к числовому значению")
			} else {
				fmt.Printf("Число: %d\n", a)
			}
		}
	}

	fmt.Println("------------------------------------------------")
}
