package task02

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zoosmand/usecons/v3"
)

const threshold = 25

func Returns() {
	h := "Задача 14.2. Функция, возвращающая несколько значений."
	fmt.Print(usecons.Header(&h))

	x, y := coordGen()
	fmt.Printf("Инициируемые значения: %d, %d\n", x, y)

	doFormula(&x, &y)
	fmt.Printf("Результат обработки: %d, %d\n\n", x, y)
}

func coordGen() (int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(threshold), r.Intn(threshold)
}

func doFormula(x *int, y *int) {
	*x = *x*2 + 10
	*y = *y*(-3) - 5
}
