package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	pieces   = 6
	expected = []int{1, 1, 2, 2, 2, 8}
)

// https://open.kattis.com/contests/dasomz/problems/bijele
func main() {
	in.Split(bufio.ScanWords)

	given := make([]int, pieces)
	for i := 0; i < pieces; i++ {
		given[i] = nextInt()
	}

	for _, v := range validate(given) {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func validate(given []int) []int {
	diff := make([]int, pieces)
	for i := 0; i < pieces; i++ {
		diff[i] = expected[i] - given[i]
	}
	return diff
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
