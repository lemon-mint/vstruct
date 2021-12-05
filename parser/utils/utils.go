package utils

import "fmt"

func PrintFileError(fileName string, line, column int, data string) {
	fmt.Printf("%s:%d:%d: %s\n", fileName, line, column, data)
}

func Expected(expected interface{}, data interface{}) string {
	return fmt.Sprintf("Expected %v, but found %v\n", expected, data)
}
