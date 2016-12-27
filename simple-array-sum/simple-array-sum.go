package main

import (
	"fmt"
	"github.com/evansgp/hackerrank-go/hackerutil"
)

func main() {
	foo := hackerutil.ReadInts()
	var sum = Sum(foo)
	fmt.Println(sum)
}

func Sum(xs []int) int {
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum
}
