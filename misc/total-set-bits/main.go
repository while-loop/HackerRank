package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

// https://www.geeksforgeeks.org/count-total-set-bits-in-all-numbers-from-1-to-n/
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(CountSetBits(NextUint()))
}

func CountSetBits(n uint) int {

	count := 0
	for i := uint(0); i <= n; i++ {
		count += bits.OnesCount(i)
	}

	return count
}

func NextUint() uint {
	in.Scan()
	tmp, err := strconv.ParseUint(in.Text(), 10, 32)
	if err != nil {
		panic(err)
	}

	return uint(tmp)
}
