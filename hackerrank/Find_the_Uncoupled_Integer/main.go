/**
  * XOR each number to cancel out (0x00) for paired numbers
  */

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

/**
1, 2, 3, 1, 2
3

1, 2, 3, 4, 5, 99, 1, 2, 3, 4, 5
99
 */
func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	fmt.Println(uncoupled(s))
}

func uncoupled(s string) int {
	tmp := 0
	for _, element := range strings.Split(s, ",") {
		number, err := strconv.Atoi(strings.TrimSpace(element))
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			break
		}
		tmp = tmp ^ number
	}
	return tmp
}