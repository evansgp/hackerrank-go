package main

import (
	"fmt"
)

func ExampleSolve_simple() {
	fmt.Println(solve([]int{1, 2}))
	// Output:
	// 3
}

func ExampleSolve_empty() {
	fmt.Println(solve([]int{}))
	// Output:
	// 0
}

func ExampleSolve_zeros() {
	fmt.Println(solve([]int{0, 0}))
	// Output:
	// 0
}

func ExampleSolve_negatives() {
	fmt.Println(solve([]int{2, -1}))
	// Output:
	// 1
}

func ExampleSolve_all_negatives() {
	fmt.Println(solve([]int{-2, -1}))
	// Output:
	// -3
}