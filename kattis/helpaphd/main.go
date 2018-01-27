package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	np      = "P=NP"
	skipped = "skipped"
)

// https://open.kattis.com/contests/bfe8oz/problems/helpaphd
func main() {
	in.Split(bufio.ScanLines)

	N := nextInt()
	for i := 0; i < N; i++ {
		fmt.Println(help(nextLine()))
	}
}

func help(s string) string {
	if s == np {
		return skipped
	}

	tokens := strings.Split(s, "+")
	return strconv.Itoa(i(tokens[0]) + i(tokens[1]))

}

func i(a string) int {
	tmp, _ := strconv.Atoi(a)
	return tmp
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}

func nextLine() string {
	in.Scan()
	return in.Text()
}
