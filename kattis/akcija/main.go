package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// https://open.kattis.com/contests/dasomz/problems/akcija
func main() {
	in.Split(bufio.ScanWords)

	numBooks := nextInt()
	books := make([]int, numBooks)
	for i := 0; i < numBooks; i++ {
		books[i] = nextInt()
	}
	fmt.Println(totalCost(books))
}

func totalCost(books []int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(books)))
	ttl := int64(0)
	for i := 0; i < len(books); i++ {
		if (i+1)%3 == 0 {
			continue
		}

		ttl += int64(books[i])
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
