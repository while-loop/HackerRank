package main

import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

type Shield struct {
	lower, upper float64
	factor       float64
}

func init() {
	in.Split(bufio.ScanWords)
}

// https://open.kattis.com/contests/bw8c4i/problems/carefulascent
func main() {
	x, y, sLen := nextFloat64(), nextFloat64(), nextInt()
	shields := make([]Shield, sLen)
	for i := 0; i < sLen; i++ {
		shields[i] = Shield{nextFloat64(), nextFloat64(), nextFloat64()}
	}

	fmt.Printf("%.11f\n", horizontalVelocity(x, y, shields))
}

func horizontalVelocity(x, y float64, shields []Shield) float64 {
	traveledX := float64(0)
	for _, s := range shields {
		traveledX += (s.upper - s.lower) * s.factor
		y -= s.upper - s.lower
	}

	return x / (traveledX + y)
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
