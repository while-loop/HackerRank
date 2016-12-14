/**
  * Basic LIFO Stack
  *
  */

package main

import (
	"fmt"
	"errors"
	"bufio"
	"os"
	"strings"
)

/**
12
144

30
832040
 */

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	fmt.Println(IsBalanced(strings.TrimSpace(s)))
}

func IsBalanced(s string) string {
	stack := Stack{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack.Push(byte(r))
		} else if r == ')' || r == ']' || r == '}' {
			val, err := stack.Pop()
			if err != nil {
				return "False"
			}

			if r == ')' && val != '(' {
				return "False"
			} else if r == ']' && val != '[' {
				return "False"
			} else if r == '}' && val != '{' {
				return "False"
			}
		}
	}

	if stack.Empty() {
		return "True"
	} else {
		return "False"
	}
}

type Stack struct {
	elements []byte
	size     int
}

func (this *Stack) Empty() bool {
	return this.size == 0
}

func (this *Stack) Peek() (byte, error) {
	if this.Empty() {
		return 0, errors.New("Empty Stack")
	} else {
		return this.elements[0], nil
	}
}

func (this *Stack) Pop() (byte, error) {
	if this.Empty() {
		return 0, errors.New("Empty Stack")
	} else {
		this.size-- // decrement before return because of 0-index
		ele := this.elements[this.size]
		this.elements[this.size] = 0
		return ele, nil
	}
}

func (this *Stack) Push(val byte) {
	if this.size >= len(this.elements) {
		// resize elements array
		eleArrSize := len(this.elements)
		if eleArrSize == 0 {
			eleArrSize = 1
		}
		tmp := make([]byte, eleArrSize * 2)
		copy(tmp, this.elements)
		this.elements = tmp
	}

	this.elements[this.size] = val
	this.size++
}

func (this *Stack) Size() int {
	return this.size
}