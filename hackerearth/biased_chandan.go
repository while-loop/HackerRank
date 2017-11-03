package main

import (
	"../collection/go"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/biased-chandan/
 */
func cv() {
	fmt.Print(SolveBias(bufio.NewScanner(os.Stdin)))
}

func SolveBias(scanner *bufio.Scanner) int {
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	nums, _ := strconv.Atoi(scanner.Text())
	stack := collection.Stack{}
	var tmp int
	for i := 0; i < nums; i++ {
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())

		if tmp == 0 {
			stack.Pop()
		} else {
			stack.Push(tmp)
		}

	}

	return sum(stack)
}

func sum(s Stack) int {
	sum := 0
	for !s.Empty() {
		a, _ := s.Pop()
		sum += a.(int)
	}
	return sum
}
