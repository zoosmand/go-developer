package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Introduction
	fmt.Println()
	fmt.Print("Задача 8.3. Лимонад киоск\n\n")

	allowedBills := [3]int{5, 10, 20}

	// ----------------------------------------------------
	// Cast string live looks like "5,10,20" into int array
	// and examine each value with the allowed ones
	var examinBillSerie = func(serie string) (bool, []int) {
		result := true
		var bills []int
		// bills = bills[:0]
		billValues := strings.Split(serie, ",")

		for i := 0; i < len(billValues); i++ {
			_billValue, err := strconv.Atoi(billValues[i])

			if err != nil {
				result = false
			}

			checkFlag := 0
			for j := 0; j < len(allowedBills); j++ {

				if _billValue == allowedBills[j] {
					checkFlag |= 1
					bills = append(bills, _billValue)
					break
				}
			}

			if checkFlag != 1 {
				result = false
			}
		}
		return result, bills
	}

	// ----------------------------------------------------
	// Sum of values of an integer array given as an input parameter
	var arraySum = func(arr []int) int {
		sum := 0
		for i := 0; i < len(arr); i++ {
			sum += arr[i]
		}
		return sum
	}

	// ----------------------------------------------------
	// Find change and print each sale report
	// Input parameter gives an array with bills
	// Each bill presumes an expecting sale
	var lemonadeChange = func(bills []int) (bool, int, []int) {
		// The result of searching change
		var result bool
		// Successfull sale indicator
		var successfullSale bool
		// Bills in cashier after each operation
		var billsInCashier []int
		// Sales counter
		salesCount := 0

		for i := 0; i < len(bills); i++ {
			result = false
			successfullSale = false
			switch bills[i] {
			// when $5 bill appeears, then no change back, sale is done
			case 5:
				billsInCashier = append(billsInCashier, bills[i])
				salesCount++
				successfullSale = true
			// when $10 bill appeears, then look for only $5 bill in cashier for change
			case 10:
				for j := 0; j < len(billsInCashier); j++ {
					if billsInCashier[j] == 5 {
						// when $5 bill has found, remove it from cachier, then add $10 bill to the cachier
						billsInCashier[len(billsInCashier)-1], billsInCashier[j] = billsInCashier[j], billsInCashier[len(billsInCashier)-1]
						billsInCashier = billsInCashier[:len(billsInCashier)-1]
						billsInCashier = append(billsInCashier, bills[i])
						result = true
						salesCount++
						successfullSale = true
						break
					}
				}
			// when $20 bill appeears,
			// first, check if enought money in the cachier
			// then, collect bill with two conditions: 5+5+5 or 5+10
			// when change is collected, recalculate bills in the cashier
			case 20:
				//  collect indexes of potential bills for a change
				var changeBillIndexes []int
				change := 0
				if arraySum(billsInCashier) >= 15 {
				fineChangeLoop:
					for i := 0; i < len(billsInCashier); i++ {
						switch change {
						case 0, 5:
							if billsInCashier[i] == 5 || billsInCashier[i] == 10 {
								change += billsInCashier[i]
								changeBillIndexes = append(changeBillIndexes, i)
							}
						case 10:
							if billsInCashier[i] == 5 {
								change += billsInCashier[i]
								changeBillIndexes = append(changeBillIndexes, i)
								// this avoids a situation when three $5 bills in the cashier only
								result = true
								salesCount++
								successfullSale = true
							}
						case 15:
							result = true
							salesCount++
							successfullSale = true
							break fineChangeLoop
						default:
							result = false

						}
					}
					if result {
						// zero indexed biils
						for i := 0; i < len(changeBillIndexes); i++ {
							for j := 0; j < len(billsInCashier); j++ {
								if j == changeBillIndexes[i] {
									billsInCashier[j] = 0
									break
								}
							}
						}
						// remove zeroed bills via temporary array
						var newbillsInCashier []int
						for i := 0; i < len(billsInCashier); i++ {
							if billsInCashier[i] != 0 {
								newbillsInCashier = append(newbillsInCashier, billsInCashier[i])
							}
						}
						billsInCashier = newbillsInCashier
						billsInCashier = append(billsInCashier, bills[i])
					}
				}
			}
			fmt.Print("Предложена купюра: ", bills[i], ", Удачная продажа: ", successfullSale, ", Сдача дана: ", result, ", Купюр в кассе: ", billsInCashier, "\n")
		}
		return result, salesCount, billsInCashier
	}

	var billSerie string
	fmt.Println("Введите серию купюр, например: 5,10,20")

	for {
		_, _ = fmt.Scan(&billSerie)

		checkResult, bills := examinBillSerie(billSerie)
		if checkResult {
			fmt.Println("---------------------")
			_, totalSales, billsInCashier := lemonadeChange(bills)
			fmt.Println("---------------------")
			fmt.Println("Всего продаж:", totalSales, ", Купюр в кассе:", billsInCashier)
			fmt.Println("---------------------")
			break
		} else {
			fmt.Println("Неверное значение! Введите серию купюр, например: 5,10,20")
		}
	}
}
