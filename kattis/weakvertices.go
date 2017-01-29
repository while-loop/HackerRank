package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)

/**
  * https://open.kattis.com/contests/ze6qt6/problems/weakvertices
  */
func main() {
	in.Split(bufio.ScanWords)

	for {
		N := NextInt()
		if N == -1 {
			break
		}
		graph := make([][]int, N)

		for i := 0; i < N; i++ {
			graph[i] = make([]int, N)
			for j := 0; j < N; j++ {
				graph[i][j] = NextInt()
			}
		}

		adj := []int{}
		for r := 0; r < N; r++ {
			for c := 0; c < N; c++ {
				for k := 0; k < N; k++ {

				}
			}
		}
		sort.Ints(adj)
		fmt.Println(adj)
	}
}

func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp

}

func NextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp

}
