package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://open.kattis.com/contests/bfe8oz/problems/chineseremainder
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {
		fmt.Println(chineseRemainder(nextInt(), nextInt(), nextInt(), nextInt()))
	}
}

func chineseRemainder(a, n, b, m int) (int, int) {
	return 0, 0
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
