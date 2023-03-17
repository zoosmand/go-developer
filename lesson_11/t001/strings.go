package t001

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/zoosmand/usecons/v3"
)

func Strings() {

	h := "Задача 11.1. Строки."
	usecons.Header(&h)

	// ----------------------------------------------------
	myLine := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	// myLine := "     "
	// myLine := ""
	myArray := strings.Split(myLine, " ")
	counter := 0

	rule := regexp.MustCompile(`^[A-Z]`)

	for i := 0; i < len(myArray); i++ {
		if rule.MatchString(myArray[i]) {
			counter++
		}
	}

	fmt.Println("------------------------------------------------")
	fmt.Printf("Количество слов начинающихся с большой буквы - %d\n", counter)
	fmt.Println("------------------------------------------------")
}
