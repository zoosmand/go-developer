package t003

import (
	"fmt"
	"log"
	"os"

	"github.com/zoosmand/usecons/v3"
)

func IReader() {
	h := "Задача 12.3. Интерфейс io.Reader."
	fmt.Println(usecons.Header(&h))

	logFile := "log.txt"

	file, err := os.OpenFile(logFile, os.O_RDONLY, 0)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// ----------------------------------------------------
	stat, err := os.Stat(logFile)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, stat.Size())

	if _, err := file.Read(buf); err != nil {
		log.Panic(err)
	}

	// ----------------------------------------------------
	fmt.Printf("File size is %d bytes\r\n", stat.Size())
	fmt.Println("------------------------------------------------")
	fmt.Print(string(buf))
	fmt.Println("------------------------------------------------")
}
