package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

var in = bufio.NewScanner(os.Stdin)

/**
  * https://open.kattis.com/problems/rationalsequence3
  */
func main() {
	in.Split(bufio.ScanWords)

	N := NextInt()
	for i := 1; i <= N; i++ {
		NextInt()
		p, q := Solve(NextInt())
		fmt.Printf("%d %d/%d\n", i, p, q)
	}
}

func Solve(n int) (uint, uint) {
	level := int(math.Floor(math.Log2(float64(n))))
	levelVal := int(math.Pow(2, float64(level)))

	p, q := uint(1), uint(1)
	levelVal /= 2
	for ; levelVal > 0; levelVal /= 2 {
		if levelVal&n == 0 { //left
			q += p
		} else { // right
			p += q
		}
	}

	return p, q
}

func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
