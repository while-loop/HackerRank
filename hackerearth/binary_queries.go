package main

import (
	"bufio"
	"os"
	"fmt"
	"bytes"
	"strconv"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/range-query-2/description/
  */
func ge5g() {
	fmt.Print(BinaryQueries(bufio.NewScanner(os.Stdin)))
}
func BinaryQueries(scanner *bufio.Scanner) string {
	var buffer bytes.Buffer
	scanner.Split(bufio.ScanWords)

	scanner.Scan(); length, _ := strconv.Atoi(scanner.Text())
	scanner.Scan(); lines, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		scanner.Scan(); arr[i], _ = strconv.Atoi(scanner.Text())
	}

	var cmd, X, R int
	for i := 0; i < lines; i++ {
		scanner.Scan(); cmd, _ = strconv.Atoi(scanner.Text())
		if cmd == 1 {
			scanner.Scan(); X, _ = strconv.Atoi(scanner.Text())
			if arr[X - 1] == 0 {
				arr[X - 1] = 1
			} else {
				arr[X - 1] = 0
			}
		} else {
			scanner.Scan(); _, _ = strconv.Atoi(scanner.Text())
			scanner.Scan(); R, _ = strconv.Atoi(scanner.Text())
			if arr[R - 1] == 1 {
				buffer.WriteString("ODD\n")
			} else {
				buffer.WriteString("EVEN\n")
			}
		}
	}

	return buffer.String()
}