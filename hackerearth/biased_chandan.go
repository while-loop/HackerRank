package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"errors"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/biased-chandan/
  */
func cv() {
	fmt.Print(SolveBias(bufio.NewScanner(os.Stdin)))
}

func SolveBias(scanner *bufio.Scanner) int {
	scanner.Split(bufio.ScanWords)

	scanner.Scan(); nums, _ := strconv.Atoi(scanner.Text())
	stack := Stack{}
	var tmp int
	for i := 0; i < nums; i++ {
		scanner.Scan(); tmp, _ = strconv.Atoi(scanner.Text())

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
		return this.elements[0], nil
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
