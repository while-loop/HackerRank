package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
	"math"
)

func init() {
	in.Split(bufio.ScanWords)
}

// https://open.kattis.com/contests/naipc18-p05/problems/judgingmoose
func main() {
	fmt.Println(judge(nextInt(), nextInt()))

}
func judge(a int, b int) string {
	if a == 0 && b == 0 {
		return "Not a moose"
	} else if a == b {
		return fmt.Sprint("Even ", a+b)
	} else {
		return fmt.Sprint("Odd ", int(math.Max(float64(a), float64(b)))*2)
	}
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
