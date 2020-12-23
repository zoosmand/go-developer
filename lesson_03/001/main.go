package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 1. Конструктор квестов.")
	fmt.Println()
	
	var planetName string
	var starSystemName string
	var rangerName string
	var awardSum int
	
	fmt.Println()
	fmt.Println("Введите название планеты:")
	fmt.Scan(&planetName)
	
	fmt.Println("Введите имя звездной системы:")
	fmt.Scan(&starSystemName)
	
	fmt.Println("Введите рейнджера:")
	fmt.Scan(&rangerName)
	
	fmt.Println("Введите сумму вознаграждения:")
	fmt.Scan(&awardSum)
	
	fmt.Println()
	fmt.Println()
	fmt.Print("Как вам, ", rangerName, ", известно, мы - раса мирная, поэтому на наших военных кораблях используются",
		"наемники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли",
		"людей с планеты ", planetName, " системы ", starSystemName, ".\n\n",
		"Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал,",
		"и теперь не так-то просто завербовать призывников. Главный комиссар планеты ", planetName, ", впрочем, отлично",
		"справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало!",
		"Призыв под угрозой срыва!\n\n",
		rangerName, ", вы должны прибыть на планету ", planetName, " системы ", starSystemName, " и помочь выполнить",
		"план призыва. Мы готовы выплатить вам премию в ", awardSum, " кредитов за эту маленькую услугу.\n\n")
}
