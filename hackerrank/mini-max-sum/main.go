package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/**
 * https://www.hackerrank.com/challenges/mini-max-sum/problem
 */
func main() {
	in.Split(bufio.ScanWords)

	sum, mini, maxi := uint64(0), uint64(math.MaxUint64), uint64(0)

	for i := 0; i < 5; i++ {
		num := nextUint64()
		sum += num
		mini = min(mini, num)
		maxi = max(maxi, num)
	}

	fmt.Println(sum-maxi, sum-mini)
}

var in = bufio.NewScanner(os.Stdin)

func nextUint64() uint64 {
	in.Scan()
	tmp, _ := strconv.ParseUint(in.Text(), 10, 64)
	return tmp
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}
