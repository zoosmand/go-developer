package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 2. Симулятор маршрутки.")
	fmt.Println()

	var totalPassengers int
	var outPassengers int
	var inPassengers int
	var totalTrips int

	fmt.Println("Прибываем на остановку Улица Программистов. В салоне пассажиров:", totalPassengers)
	fmt.Println()

	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Println(0)

	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&inPassengers)
	totalPassengers += inPassengers
	totalTrips += inPassengers

	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", totalPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println()

	/* ---------------------------------------------------------------------------------------------- */
	fmt.Println("Прибываем на остановку Улица Алгоритмов. В салоне пассажиров:", totalPassengers)
	fmt.Println()

	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&outPassengers)
	totalPassengers -= outPassengers

	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&inPassengers)
	totalPassengers += inPassengers
	totalTrips += inPassengers

	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", totalPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println()

	/* ---------------------------------------------------------------------------------------------- */
	fmt.Println("Прибываем на остановку Улица Рабочих Станций. В салоне пассажиров:", totalPassengers)
	fmt.Println()

	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&outPassengers)
	totalPassengers -= outPassengers

	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&inPassengers)
	totalPassengers += inPassengers
	totalTrips += inPassengers

	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", totalPassengers)
	fmt.Println("-----------Едем---------")
	fmt.Println()

	/* ---------------------------------------------------------------------------------------------- */
	fmt.Println("Прибываем на остановку Улица Серверная. В салоне пассажиров:", totalPassengers)
	fmt.Println()

	fmt.Println("Сколько пассажиров вышло на остановке?")
	fmt.Scan(&outPassengers)
	totalPassengers -= outPassengers

	fmt.Println("Сколько пассажиров зашло на остановке?")
	fmt.Scan(&inPassengers)
	totalPassengers += inPassengers
	totalTrips += inPassengers

	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", totalPassengers)
	fmt.Println("-----------Маршрут закончен---------")
	fmt.Println()
	fmt.Println()

	tripCost := 20
	totalIncome := totalTrips * tripCost
	fmt.Print("Всего заработали: ", totalIncome, " руб.\n")

	driverSalary := totalIncome / 4
	fmt.Print("Зарплата водителя: ", driverSalary, " руб.\n")

	fuelExpenses := totalIncome / 5
	fmt.Print("Расходы на топливо: ", fuelExpenses, " руб.\n")

	taxes := totalIncome / 5
	fmt.Print("Налоги: ", taxes, " руб.\n")

	carReparing := totalIncome / 5
	fmt.Print("Расходы на ремонт машины: ", carReparing, " руб.\n")
	fmt.Println()

	totalProfit := totalIncome - driverSalary - fuelExpenses - taxes - carReparing
	fmt.Print("Итого доход: ", totalProfit, " руб.\n")
	fmt.Println()
}
