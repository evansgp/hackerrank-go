package main

import (
	"fmt"
)

func main() {
	foo := readInts()
	var sum = solve(foo)
	fmt.Println(sum)
}

func solve(xs []int) int {
	return nil
}

func readInts() []int {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}
	all := make([]int, n)
	for i := range all {
		if _, err := fmt.Scan(&all[i]); err != nil {
			panic(err)
		}
	}

	return all
}