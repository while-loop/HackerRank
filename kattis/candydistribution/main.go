package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	impossible = "IMPOSSIBLE"
)

// https://open.kattis.com/contests/bfe8oz/problems/candydistribution
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {
		d := candyDistribtion(nextInt(), nextInt())
		if d > 0 {
			fmt.Println(d)
		} else {
			fmt.Println(impossible)
		}
	}
}

func candyDistribtion(k, c int) int {
	if k%c == 0 {

	}
	return -1
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
