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

// https://open.kattis.com/contests/bw8c4i/problems/digitsum
func main() {
	cases := nextInt()
	for i := 0; i < cases; i++ {
		fmt.Println(sum(nextInt64(), nextInt64()))
	}
}

func sum(low, high int64) int64 {
	ttl := int64(0)
	for i := low; i <= high; i++ {
		ttl += digitSum(i)
	}

	return ttl
}
func digitSum(i int64) int64 {
	s := int64(0)
	for ; i > 0; i /= 10 {
		s += i%10
	}

	return s
}

var in = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}

	return int64(tmp)
}

func nextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}

	return tmp
}

