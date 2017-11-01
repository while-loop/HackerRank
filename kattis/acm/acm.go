package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var _in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/acm
 */
func main() {
	_in.Split(bufio.ScanWords)
	c, t := solve(_in)
	fmt.Println(c, t)
}

func solve(in *bufio.Scanner) (int, int) {
	penalties := map[string]int{}
	right := 0
	ttl := 0
	for {
		t := nextInt(in)
		if t == -1 {
			break
		}

		prob, answer := next(in), next(in)
		if answer == "wrong" {
			penalties[prob]++
		} else {
			ttl += t + (penalties[prob] * 20)
			right++
		}
	}

	return right, ttl
}

func next(in *bufio.Scanner) string {
	in.Scan()
	return in.Text()
}

func nextInt(in *bufio.Scanner) int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
