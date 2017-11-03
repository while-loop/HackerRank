/**
 * Iterative Factorial w/o memorization
 *
 */

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/**
3
6

0
1

7
5040

10
3628800
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	fmt.Printf("%d\n", factorial(n))
}

func factorial(n int) *big.Int {
	ret := big.NewInt(1)

	for i := 1; i <= n; i++ {
		ret.Mul(ret, big.NewInt(int64(i)))
	}

	return ret
}
