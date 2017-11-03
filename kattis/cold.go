package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/cold
 */
func main65() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	N, _ := strconv.Atoi(in.Text())
	count := 0
	for i := 0; i < N; i++ {
		in.Scan()
		tmp, _ := strconv.Atoi(in.Text())
		if tmp < 0 {
			count++
		}
	}

	fmt.Println(count)
}
