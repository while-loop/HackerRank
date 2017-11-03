package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/the-amazing-race-1/
 */
func main54() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	tests, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < tests; i++ {
		scanner.Scan()
		nums, _ := strconv.Atoi(scanner.Text())
		vals := make([]int, nums)
		for j := 0; j < nums; j++ {
			scanner.Scan()
			vals[j], _ = strconv.Atoi(scanner.Text())
		}
		maxIdx := [2]int{-1, -1} // hold max index and max value of car
		for j := 0; j < nums; j++ {
			tmp := MaxSight(vals, j)
			if tmp > maxIdx[1] {
				maxIdx[0] = j
				maxIdx[1] = tmp
			}
		}
		fmt.Println(maxIdx[0])
	}
}

func MaxSight(arr []int, pos int) int {
	count := 0
	arrLen := len(arr)
	currHeight := arr[pos]
	front := true
	back := true

	for i := 0; i < arrLen; i++ {
		tmpHeight := arr[i]
		if ((pos - i) >= 0) && front && currHeight >= tmpHeight { // cars in front
			count++
			if tmpHeight >= currHeight && (pos != i) {
				front = false
			}
		}

		if (pos+i) < arrLen && back && currHeight >= tmpHeight { // cars behind
			count++
			if tmpHeight >= currHeight && (pos != i) {
				back = false
			}
		}
	}

	return count * (pos + 1)
}
