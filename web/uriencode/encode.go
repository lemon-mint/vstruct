package main

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	var mimetype string = os.Args[1]
	var inputFileName string = os.Args[2]
	var outputFileName string = os.Args[3]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	var bufOutFile = bufio.NewWriter(outputFile)
	bufOutFile.WriteString("data:" + mimetype + ";base64,")

	var b64out = base64.NewEncoder(base64.RawURLEncoding, bufOutFile)
	_, err = io.Copy(b64out, inputFile)
	if err != nil {
		panic(err)
	}
	b64out.Close()
}
