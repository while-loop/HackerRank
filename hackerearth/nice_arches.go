package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/while-loop/competitive-programming/collection/go"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/nice-arches-1/
 */
func main3() {
	// use Reader instead of Scanner cause scanner's max string length is 2^16, our max input len is 10^5
	reader := bufio.NewReader(os.Stdin)

	tmp, _ := reader.ReadString('\n')
	tests, _ := strconv.Atoi(strings.TrimSpace(tmp))
	count := 0
	for i := 0; i < tests; i++ {
		word, _ := reader.ReadString('\n')
		if isBubbly(strings.TrimSpace(word)) {
			count++
		}
	}

	fmt.Println(count)
}

func isBubbly(s string) bool {
	stack := collection.Stack{}
	for _, value := range s {
		if stack.Empty() {
			stack.Push(value)
		} else {
			peeked, _ := stack.Peek()
			if peeked == value {
				stack.Pop()
			} else {
				stack.Push(value)
			}
		}
	}

	return stack.Empty()
}
