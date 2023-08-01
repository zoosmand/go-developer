package task02

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func AnonymousFunctions() {
	h := "Task 21.2 Anonymous functions"
	fmt.Println(usecons.Header(&h))

	a := 7
	b := 9
	fmt.Printf("%d * %d = %d\n", a, b, CallbackExecutor(func(a int, b int) int { return a * b }, a, b))
	fmt.Printf("%d - %d = %d\n", a, b, CallbackExecutor(func(a int, b int) int { return a - b }, a, b))
	fmt.Printf("%d + %d = %d\n", a, b, CallbackExecutor(func(a int, b int) int { return a + b }, a, b))

}

func CallbackExecutor(f func(int, int) int, a int, b int) int {
	return f(a, b)
}
