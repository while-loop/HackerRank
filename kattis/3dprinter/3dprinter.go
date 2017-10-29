package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/3dprinter
 */
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(solve(nextInt()))
}

func solve(statuses int) int {
	return int(math.Ceil(math.Log2(float64(statuses)))) + 1
}

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
