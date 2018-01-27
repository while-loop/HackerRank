package main

import (
	"bufio"
	"os"
	"strconv"
)

//
func main() {
	in.Split(bufio.ScanWords)

	N := nextInt()
	for i := 0; i < N; i++ {

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

func nextFloat64() float64 {
	in.Scan()
	tmp, err := strconv.ParseFloat(in.Text(), 64)
	if err != nil {
		panic(err)
	}
	return tmp
}

func nextLine() string {
	in.Scan()
	return in.Text()
}

func next() string {
	in.Scan()
	return in.Text()
}
