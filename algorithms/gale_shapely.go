package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var in = bufio.NewScanner(os.Stdin)

/**
  *
  */
func main() {
	in.Split(bufio.ScanWords)

	N := NextInt()
	for i := 0; i < N; i++ {

	}

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
