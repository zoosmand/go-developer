package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Println("Задача 4.6. Студенты")
	fmt.Println()

	var studentCount int
	fmt.Println("Введите количество студентов:")
	fmt.Scan(&studentCount)

	var groupCount int
	fmt.Println("Введите количество групп:")
	fmt.Scan(&groupCount)

	var studentNumber int
	fmt.Println("Введите номер студента:")
	fmt.Scan(&studentNumber)

	// check the student's number doesn't exceed count of students
	if studentNumber <= studentCount {
		// check the group count to be not less than 2 students per group
		if groupCount < (studentCount / 2) {
			// calculate students per group number
			studentsInGroup := studentCount / groupCount
			if (studentCount % groupCount) > 0 {
				studentsInGroup++
			}

			// calculate student in grop
			studentsGroup := studentNumber / studentsInGroup
			if (studentNumber % studentsInGroup) > 0 {
				studentsGroup++
			}

			fmt.Println()
			fmt.Print("Студент с номером ", studentNumber, " зачислен в группу ", studentsGroup, ".\n")
			fmt.Println()
		} else {
			// exception explanation if a group number exceeds student's count divided into two
			fmt.Println()
			fmt.Print("В одной группе должно быть по меньшей мере 2 студента. Введите данные заново.\n")
			fmt.Println()
		}
	} else {
		// exception explanation if a student number exceeds student's count
		fmt.Println()
		fmt.Print("Номер студента ", studentNumber, " должен быть меньше или равен количеству студентов. Введите данные заново.\n")
		fmt.Println()
	}
}
