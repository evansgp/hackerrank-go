package hackerutil

import "fmt"

func ReadInts() []int {
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
