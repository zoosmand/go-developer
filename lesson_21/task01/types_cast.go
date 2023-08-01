package task01

import (
	"fmt"

	"github.com/zoosmand/usecons/v3"
)

func Intro() {
	h := "Task 21.1 Type Casting"
	fmt.Println(usecons.Header(&h))

	fmt.Printf("%f\n\n", TypesCasting(32767, 255, 43.222))
}

func TypesCasting(x int16, y uint8, z float32) float32 {
	return (float32)(2*x) + (float32)(y^2) - 3/z
}
