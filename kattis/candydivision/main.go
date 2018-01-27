package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

//https://open.kattis.com/contests/bfe8oz/problems/candydivision
func main() {
	in.Split(bufio.ScanWords)
	c := nextInt()
	h := factors(c)
	if len(h) > 0 {
		fmt.Print((c / h[0]) - 1)
	}
	for i := 1; i < len(h); i++ {
		fmt.Print(" ", (c/h[i])-1)
	}
}

func factors(n int) []int {
	o := map[int]bool{}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 1; i <= max; i++ {
		if n%i != 0 {
			continue
		}

		o[i] = true
		o[n/i] = true
	}

	a := make([]int, len(o))
	i := 0
	for key := range o {
		a[i] = key
		i++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a
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
