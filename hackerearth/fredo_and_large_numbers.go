package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"bytes"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/long-atm-queue-3/
  */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var buffer bytes.Buffer

	scanner.Scan();
	N, _ := strconv.Atoi(scanner.Text())

	nums := map[int]int{}
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan(); arr[i], _ = strconv.Atoi(scanner.Text())
		nums[arr[i]]++
	}

	scanner.Scan();
	N, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < N; i++ {
		scanner.Scan(); qry, _ := strconv.Atoi(scanner.Text())
		scanner.Scan(); num, _ := strconv.Atoi(scanner.Text())
		buffer.WriteString(fmt.Sprintf("%d\n", freq(nums, arr, num, qry)))
	}

	fmt.Print(buffer.String())
}

func freq(freqs map[int]int, arr []int, freq int, exact int) int {

	for i := 0; i < len(arr); i++ {
		if (exact == 0) && (freqs[arr[i]] >= freq) {
			return arr[i]
		} else if (exact == 1) && (freqs[arr[i]] == freq) {
			return arr[i]
		}
	}

	return 0
}
