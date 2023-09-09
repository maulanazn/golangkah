package function

import (
	"fmt"
	"strings"
)

func PrintHello(message string, arr []string) (string, string) {
	var nameString string = strings.Join(arr, " ")
	fmt.Println(message, nameString)

	return message, nameString
}

func DumpData(numbers ...int) (int, error) {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	return total, nil
}
