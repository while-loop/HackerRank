package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/long-atm-queue-3/
  */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan(); N, _ := strconv.Atoi(scanner.Text())
	cur := math.MaxInt32
	groups := 0
	for i := 0; i < N; i++ {
		scanner.Scan(); tmp, _ := strconv.Atoi(scanner.Text())
		if tmp < cur {
			groups++
		}
		cur = tmp
	}

	fmt.Println(groups)
}
