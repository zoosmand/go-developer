package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	usrcons "../../packages/usrcons"
)

func main() {
	h := "Задача 12.2. Лог файл."
	usrcons.PrintHeader(&h)

	logFile := "log.txt"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		// fmt.Printf("Error %s\r\n", err.Error())
		log.Panic(err)
	}
	defer file.Close()

	// ----------------------------------------------------
	// read last line, get last used number of log file line
	lastNumber := 0
	const BUFSIZE = 256
	const NUMLEN = 8

	var b bytes.Buffer
	buf := make([]byte, BUFSIZE)

	stat, err := os.Stat(logFile)
	if err != nil {
		log.Fatal(err)
	}

	lastEnteredNumberPos := 0
	if stat.Size() > BUFSIZE {
		start := stat.Size() - BUFSIZE

		if _, err := file.ReadAt(buf, start); err != nil {
			log.Panic(err)
		}
		b.Write(buf)

		lastEnteredNumberPos = strings.LastIndex(b.String(), "record: #")
		if lastEnteredNumberPos == -1 {
			log.Panic("Error in log file: can't find # sign")
		}

		lastEnteredNumberPos += 9

		commaPos := strings.Index(string(buf[lastEnteredNumberPos:lastEnteredNumberPos+NUMLEN]), ",")
		if commaPos == -1 {
			log.Panic("Error in log file: can't find comma")
		}

		lastNumber, err = strconv.Atoi(string(buf[lastEnteredNumberPos : lastEnteredNumberPos+commaPos]))
		if err != nil {
			log.Panic(err)
		}
	}

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
		if _, err = file.WriteString(
			fmt.Sprintf("record: #%d, %s, \"%s\"\r\n",
				lastNumber,
				time.Now().Format("2006-01-02 15:04:05 UTC-07"),
				line)); err != nil {
			log.Panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------------------------")
}
