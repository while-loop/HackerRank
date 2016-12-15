/**
  * Coin Change Dyn Prog
  *
  */

package main

import (
	"fmt"
)

/**

 */

func main() {
	var N, M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)
	Cs := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &Cs[i])
	}

	fmt.Println(MakeChange(N, M, Cs))
}

func MakeChange(N, M int, Cs []int) int {
	vals := make([]int, N + 1)

	vals[0] = 1 // base case

	for i := 0; i < M; i++ {
		for j := Cs[i]; j <= N; j++ {
			vals[j] += vals[j - Cs[i]];
		}
	}

	return vals[N]
}