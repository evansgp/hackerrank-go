package main

import (
	"fmt"
)

func ExampleSum_simple() {
	fmt.Println(Sum([]int{1, 2}))
	// Output:
	// 3
}

func ExampleSum_empty() {
	fmt.Println(Sum([]int{}))
	// Output:
	// 0
}

func ExampleSum_zeros() {
	fmt.Println(Sum([]int{0, 0}))
	// Output:
	// 0
}

func ExampleSum_negatives() {
	fmt.Println(Sum([]int{2, -1}))
	// Output:
	// 1
}

func ExampleSum_all_negatives() {
	fmt.Println(Sum([]int{-2, -1}))
	// Output:
	// -3
}