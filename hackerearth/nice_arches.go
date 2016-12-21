package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"errors"
	"strings"
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
		word , _ := reader.ReadString('\n')
		if (isBubbly(strings.TrimSpace(word))) {
			count++
		}
	}

	fmt.Println(count)
}

func isBubbly(s string) bool {
	stack := Stack{}
	for _, value := range s {
		if (stack.Empty()) {
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


/**
	Kinda generic Stack
 */
type Stack struct {
	elements []interface{}
}

func (this *Stack) Empty() bool {
	return len(this.elements) == 0
}

func (this *Stack) Peek() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Stack")
	} else {
		return this.elements[len(this.elements) - 1], nil
	}
}

func (this *Stack) Pop() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Stack")
	} else {
		size := len(this.elements) - 1
		res := this.elements[size]
		this.elements[size] = nil
		this.elements = this.elements[:size]
		return res, nil
	}
}

func (this *Stack) Push(val interface{}) {
	this.elements = append(this.elements, val)
}

func (this *Stack) Size() int {
	return len(this.elements)
}