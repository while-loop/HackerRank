package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerrank.com/challenges/compare-the-triplets/problem
 */
func main() {
	in.Split(bufio.ScanWords)

	as := []int{nextInt(), nextInt(), nextInt()}
	bs := []int{nextInt(), nextInt(), nextInt()}
	fmt.Println(cmp(as, bs))
}

func cmp(as, bs []int) (a, b int) {
	for i := 0; i < len(as); i++ {
		if as[i] > bs[i] {
			a++
		} else if as[i] < bs[i] {
			b++
		}
	}

	return
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
