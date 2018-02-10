package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

func init() {
	in.Split(bufio.ScanWords)
}

// https://open.kattis.com/contests/bw8c4i/problems/commercials
func main() {
	bLen, cost := nextInt(), nextInt()
	blocks := make([]int, bLen)

	for i := 0; i < bLen; i++ {
		blocks[i] = nextInt()
	}

	fmt.Println(maxProfit(cost, blocks))
}

func maxProfit(cost int, blocks []int) int {
	localMax, globalMax := 0, 0
	for _, v := range blocks {
		localMax = max(0, localMax+v-cost)
		globalMax = max(globalMax, localMax)
	}

	return globalMax
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}

	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}