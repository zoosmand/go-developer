package main

import (
	"fmt"
	"math/bits"
)

func main() {
	dig1 := ^-9
	dig2 := ^-103

	bb := bits.Len32(uint32(dig1))

	line := fmt.Sprintf("%b --- %b --- %d", int8(dig1), int8(dig2), bb)
	fmt.Println(line)

	line2 := fmt.Sprintf("%b --- %b", dig1, dig2)
	fmt.Println(line2)
}
