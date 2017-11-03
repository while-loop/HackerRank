package main

import (
	"./../collection/go"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monks-love-for-food/
 */
func erv() {
	fmt.Print(SolveMonkFood(bufio.NewScanner(os.Stdin)))
}

func SolveMonkFood(s *bufio.Scanner) string {
	var buffer bytes.Buffer
	scanner = s
	scanner.Split(bufio.ScanWords)

	foodPile := collection.Stack{}
	queries := readInt()
	var qtype int
	for i := 0; i < queries; i++ {
		qtype = readInt()

		if qtype == 1 {
			if foodPile.Empty() {
				buffer.WriteString("No Food\n")
			} else {
				tmp, _ := foodPile.Pop()
				buffer.WriteString(fmt.Sprintf("%d\n", tmp))
			}
		} else {
			foodPile.Push(readInt())
		}
	}

	return buffer.String()
}

func readInt() int {
	scanner.Scan()
	queries, _ := strconv.Atoi(scanner.Text())
	return queries
}
