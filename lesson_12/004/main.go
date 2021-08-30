package main

import (
	"fmt"
	"log"
	"os"

	usrcons "../../packages/usrcons"
)

func main() {
	h := "Задача 12.4. Уровни доступа."
	usrcons.PrintHeader(&h)

	logFile := "log.txt"

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR, 0444)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	aLine := "_Some data to write to the log file"
	writeLine(file, &aLine)
	fmt.Println("------------------------------------------------")
}

func writeLine(file *os.File, line *string) {
	if _, err := file.WriteString(*line); err != nil {
		log.Panic(err)
	}
}
