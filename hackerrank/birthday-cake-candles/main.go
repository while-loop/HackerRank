package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func main() {
	in.Split(bufio.ScanWords)

	nums := nextInt()

	count, max := 0, -1
	for i := 0; i < nums; i++ {
		n := nextInt()
		if n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}
	}
	fmt.Println(count)
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
