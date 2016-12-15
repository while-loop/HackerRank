/**
  * Recursive Fibonacci as per request of contest rules
  *
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
12
144

30
832040
 */
func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(s))

	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}