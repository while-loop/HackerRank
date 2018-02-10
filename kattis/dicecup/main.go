package main

import (
	"fmt"
)

// https://open.kattis.com/contests/bw8c4i/problems/dicecup
func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for _, v := range outcomes(a, b) {
		fmt.Println(v)
	}
}

func outcomes(a, b int) []int {
	out := make([]int, 0)
	if a > b {
		a, b = b, a
	}

	for i := a + 1; i < b+2; i++ {
		out = append(out, i)
	}

	return out
}
