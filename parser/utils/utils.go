package utils

import "fmt"

func PrintFileError(fileName string, line, column int, data string) {
	fmt.Printf("%s:%d:%d: %s\n", fileName, line, column, data)
}
