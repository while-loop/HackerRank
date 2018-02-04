package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/contests/dasomz/problems/mosquito
 */
func main() {
	in.Split(bufio.ScanWords)

	for in.Scan() {
		mos, pup, larv := nInt(), nextInt(), nextInt()
		eggs, rform, sAdult := nextInt(), nextInt(), nextInt()
		n := nextInt()

		fmt.Println(solve(mos, pup, larv, eggs, rform, sAdult, n))
	}
}

func solve(mos, pup, larv, eggs, rform, sAdult, n int) int {
	for i := 0; i < n; i++ {
		mos, pup, larv = pup/sAdult, larv/rform, eggs*mos
	}
	return mos
}

func nextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}
	return tmp
}

func nInt() int {
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}
	return tmp
}
