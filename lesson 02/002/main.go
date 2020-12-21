package main

import (
	"fmt"
)

func main() {
	var commodityName string = "Elecric Scooter"
	var commodityValue int = 6400
	var deliveryFee int = 350
	var standardDiscount int = 700

	fmt.Println("Товар:", commodityName)
	fmt.Println("Стоимость товара:", commodityValue)
	fmt.Println("Стоимость доставки:", deliveryFee)
	fmt.Println("Размер скидки:", standardDiscount)
	fmt.Println("---------")

	fmt.Println("Итого:", commodityValue+deliveryFee-standardDiscount)
}
