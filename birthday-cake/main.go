package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	candles := int32(0)
	for i := 0; i < len(ar); i++ {
		if i == 0 {
			candles = 1
		} else if ar[i-1] == ar[i] {
			candles++
		}
	}
	return candles
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create("/Users/jefri/Documents/workspace/go-algorithms/birthday-cake/output.file")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " "

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
