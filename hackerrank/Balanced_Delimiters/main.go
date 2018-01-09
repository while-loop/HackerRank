/**
 * Basic LIFO Stack
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/while-loop/competitive-programming/collection/go"
)

/**
()[]{}
([{}])
([]{})

([)]
([]
[])
([})
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
