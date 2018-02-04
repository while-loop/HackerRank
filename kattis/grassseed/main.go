package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type rect struct {
	l, w float64
}

func (r rect) area() float64 {
	return r.l * r.w
}

// https://open.kattis.com/contests/dasomz/problems/grassseed
func main() {
	in.Split(bufio.ScanWords)

	cost := nextFloat64()
	lawns := nextInt()
	areas := make([]rect, lawns)
	for i := 0; i < lawns; i++ {
		areas[i] = rect{w: nextFloat64(), l: nextFloat64()}
	}

	fmt.Printf("%.7f\n", totalCost(cost, areas))
}

func totalCost(cost float64, areas []rect) float64 {
	ttl := float64(0)
	for _, ar := range areas {
		ttl += cost * ar.area()
	}
	return ttl
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
