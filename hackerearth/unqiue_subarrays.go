package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"bytes"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/unique-subarrays/
  */
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var buffer bytes.Buffer

	scanner.Scan(); tests, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < tests; i++ {
		scanner.Scan(); nums, _ := strconv.Atoi(scanner.Text())
		arr := make([]int, nums)
		for j := 0; j < nums; j++ {
			scanner.Scan(); arr[j], _ = strconv.Atoi(scanner.Text())
		}

		buffer.WriteString(fmt.Sprintf("%d\n", Weights(arr)))
	}

	fmt.Print(buffer.String())
}

func Weights(arr []int) int {
	return 3
}
