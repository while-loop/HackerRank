package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/grading/problem
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {
		fmt.Println(round(nextInt()))
	}
}

func round(g int) int {
	if g < 38 {
		return g
	}

	m := g % 5
	if m > 2 {
		g += 5 - m
	}

	return g
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
