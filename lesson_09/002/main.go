package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

func main() {
	// Introduction
	fmt.Print("\nЗадача 9.2. Переполнение при умножении.\n")
	fmt.Print("-------------------------\n\n")

	// ----------------------------------------------------
	// Check user input assuming ingegers should be type of int16
	var checkUserInput = func(numbersLine string) (bool, []int16) {
		result := true
		var numbers []int16
		_numbers := strings.Split(numbersLine, ",")

		if len(_numbers) < 2 {
			return false, numbers
		}

		for i := 0; i < len(_numbers); i++ {
			// _number, err := strconv.Atoi(_numbers[i])
			_number, err := strconv.ParseInt(_numbers[i], 10, 16)

			if err != nil {
				result = false
			} else {
				numbers = append(numbers, int16(_number))
			}
		}
		return result, numbers
	}

	// ----------------------------------------------------
	// Text the integer type
	var getIntType = func(number int32, module int32, size int) string {
		if (module >> (size - 1)) == 1 {
			if number < 0 {
				return fmt.Sprintf("int%d", size*2)
			} else {
				return fmt.Sprintf("uint%d", size)
			}
		}
		return fmt.Sprintf("int%d", size)
	}

	// ----------------------------------------------------
	// Calculate necessary size of an integer value
	// resulting as a multiplication of the first two items of the array
	var checkIntSize = func(numbers []int16) string {
		multipliedNumber := int32(numbers[0]) * int32(numbers[1])
		var typeName string
		var moduledNumber int32

		if multipliedNumber < 0 {
			moduledNumber = ^multipliedNumber + 1
		} else {
			moduledNumber = multipliedNumber
		}

		length := bits.Len32(uint32(moduledNumber))

		if length <= 8 {
			length = 8
		} else if length <= 16 {
			length = 16
		} else if length <= 32 {
			length = 32
		}

		typeName = getIntType(multipliedNumber, moduledNumber, length)

		return fmt.Sprintf("Достаточный тип для сохранения результата: %s", typeName)
	}

	// ----------------------------------------------------
	var userNumbers string
	invitationLine := "Введите два числа, например: 1,-1 или: 640,100"
	fmt.Println(invitationLine)

	for {
		_, _ = fmt.Scan(&userNumbers)

		checkResult, numbers := checkUserInput(userNumbers)
		if checkResult {
			fmt.Println("---------------------")
			fmt.Println(checkIntSize(numbers))
			break
		} else {
			fmt.Println("Неверный ввод!", invitationLine)
		}
	}
}
