package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/contests/ze6qt6/problems/trollhunt
 */
func main() {
	in.Split(bufio.ScanWords)

	b := NextInt()
	k := NextInt()
	g := NextInt()

	fmt.Println(N)
}

func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}

func NextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp
}

func NextLine() string {
	in.Scan()
	return in.Text()
}

func Next() string {
	in.Scan()
	return in.Text()
}
