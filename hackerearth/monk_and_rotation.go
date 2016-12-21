package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"bytes"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-rotation-3/
  */
func main34() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var buffer bytes.Buffer

	scanner.Scan(); tests, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < tests; i++ {
		scanner.Scan(); nums, _ := strconv.Atoi(scanner.Text())
		scanner.Scan(); shifts, _ := strconv.Atoi(scanner.Text())
		arr := make([]int, nums)

		for j := 0; j < nums; j++ {
			scanner.Scan(); arr[j], _ = strconv.Atoi(scanner.Text())
		}

		buffer.WriteString(SolveRotation(arr, shifts) + "\n")
	}

	fmt.Print(buffer.String())
}

func SolveRotation(arr []int, shifts int) string {
	var buffer bytes.Buffer
	arrLen := len(arr)
	if shifts >= arrLen {
		shifts %= arrLen;
	}
	for i := 0; i < arrLen; i++ {
		pos := ((arrLen - shifts + i) % arrLen)
		buffer.WriteString(fmt.Sprintf("%d ", arr[pos]))
	}
	return buffer.String()
}