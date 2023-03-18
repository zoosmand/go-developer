package task01

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zoosmand/usecons/v3"
)

func Returns() {
	h := "Задача 14.1. Функция, возвращающая результат."
	fmt.Print(usecons.Header(&h))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := r.Int()

	fmt.Printf(" Результат обработки числа %d: %t\n\n", randInt, isEven(randInt))
}

func isEven(dig int) bool {
	if dig%2 == 0 {
		return true
	} else {
		return false
	}
}
