package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)
var store = map[string]int{}

// https://www.geeksforgeeks.org/dynamic-programming-set-8-matrix-chain-multiplication/
func main() {
	in.Split(bufio.ScanWords)

	N := NextInt()
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = NextInt()
	}

	fmt.Println(MatrixChainOrder(arr))
}

func MatrixChainOrder(matrix []int) int {
	return matrixChainOrder(matrix, 1, len(matrix)-1)
}

func matrixChainOrder(matrix []int, start, end int) int {
	if start == end { // base case
		return 0
	}

	if calcs, exists := store[fmt.Sprintf("%d%d", start, end)]; exists {
		return calcs
	}

	min := math.MaxInt32
	for i := start; i < end; i++ {
		if calcs := (matrix[start-1] * matrix[i] * matrix[end]) +
			matrixChainOrder(matrix, start, i) +
			matrixChainOrder(matrix, i+1, end); calcs < min {
			min = calcs
		}
	}

	store[fmt.Sprintf("%d%d", start, end)] = min
	return min
}

func NextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}

	return tmp
}
