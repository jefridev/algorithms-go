package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	const sign = "#"
	for i := int32(1); i <= n; i++ {
		spacesLength := n - (n - i)
		charactersLength := n - spacesLength
		for j := int32(0); j < charactersLength; j++ {
			fmt.Print(" ")
		}
		for j := int32(0); j < spacesLength; j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
