package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/r2
 */
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(solve(nextInt(), nextInt()))
}

func solve(r1, s int) int {
	return s*2 - r1
}

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
