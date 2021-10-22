package main

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func main() {
	h := "Задача 13.2. Функция, принимающая значение по ссылке."
	fmt.Print(usecons.Header(&h))

	a := 234
	b := 567
	fmt.Printf("%d %d\n\r", a, b)

	exchangeNumbers(&a, &b)
	fmt.Printf("%d %d\n\r", a, b)

}

func exchangeNumbers(a, b *int) {
	*a, *b = *b, *a
}
