package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */

	const shiftValue string = "PM"
	const morningShift string = "AM"
	pastNoon := false
	hour, err := strconv.ParseInt(s[0:2], 10, 32)
	if err != nil {
		panic(err)
	}
	shift := s[len(s)-2:]
	unmutable := s[3:8]
	if shift == shiftValue && hour != 12 {
		pastNoon = true
	}
	if shift == morningShift && hour == 12 {
		pastNoon = false
		hour = 0
	}
	if pastNoon {
		hour += 12
	}
	format := fmt.Sprintf("%02d:%s", hour, unmutable)
	return format
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
