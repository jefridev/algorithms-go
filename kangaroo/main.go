package main

import "fmt"

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x2 > x1 && v2 > v1 {
		return "NO"
	}
	for i := 0; i < 10000; i++ {
		x1, x2 = x1+v1, x2+v2
		if x1 == x2 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(kangaroo(4523, 8092, 9419, 8076))
}
