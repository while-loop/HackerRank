package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/unique-subarrays/
 */
func main32() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var buffer bytes.Buffer

	scanner.Scan()
	tests, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < tests; i++ {
		scanner.Scan()
		nums, _ := strconv.Atoi(scanner.Text())
		arr := make([]int, nums)
		for j := 0; j < nums; j++ {
			scanner.Scan()
			arr[j], _ = strconv.Atoi(scanner.Text())
		}

		buffer.WriteString(fmt.Sprintf("%d\n", Weights(arr)))
	}

	fmt.Print(buffer.String())
}

func Weights(arr []int) int {
	set := map[int]bool{} // convert to set
	for _, value := range arr {
		set[value] = true
	}

	nums := len(set)
	ttl := 0
	for i := 0; i < nums; i++ {
		ttl += (nums - i) * (i + 1)
	}

	return ttl
}
