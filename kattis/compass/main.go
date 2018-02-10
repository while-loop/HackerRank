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

// https://open.kattis.com/contests/bw8c4i/problems/compass
func main() {
	fmt.Println(position(nextInt(), nextInt()))
}

func position(current int, pos int) int {
	delta := (360 - current) + pos
	for delta > 180 {
		delta -= 360
	}

	return delta
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
