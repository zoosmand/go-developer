package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	usrcons "../../packages/usrcons"
)

func main() {
	h := "Задача 12.2. Лог файл. ioutil"
	usrcons.PrintHeader(&h)

	logFile := "./log.txt"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		// fmt.Printf("Error %s\r\n", err.Error())
		log.Panic(err)
	}
	defer file.Close()

	// ----------------------------------------------------
	lastNumber := 0
	var b bytes.Buffer
	b.Reset()

	// ----------------------------------------------------
	invite := "Введите текст объявления."
	fmt.Println(invite)

	// ----------------------------------------------------
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "выход" || line == "exit" {
			break
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		lastNumber++
		if _, err = b.WriteString(
			fmt.Sprintf("record: #%d, %s, \"%s\"\r\n",
				lastNumber,
				time.Now().Format("2006-01-02 15:04:05 UTC-07"),
				line)); err != nil {
			log.Panic(err)
		}
		if err := ioutil.WriteFile(logFile, b.Bytes(), 0644); err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------------------------")
}
