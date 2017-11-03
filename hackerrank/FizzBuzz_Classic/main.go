/**
 * FiZzBuZz
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
1
-
1

15
--
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	arr := FiZzBuZz(n)
	printFizzBuzz(arr)
}
func printFizzBuzz(arr []string) {
	for _, str := range arr {
		fmt.Println(str)
	}
}

func FiZzBuZz(n int) []string {
	arr := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			arr[i-1] = "Buzz"
		} else if i%3 == 0 {
			arr[i-1] = "Fizz"
		} else {
			arr[i-1] = strconv.Itoa(i)
		}
	}
	return arr
}
