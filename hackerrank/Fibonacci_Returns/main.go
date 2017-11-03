/**
 * Recursive Fibonacci w/ memoization
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
12
144

30
832040
*/

var dp [100]*big.Int

func main() {
	initdp()
	reader := bufio.NewReader(os.Stdin)

	for {
		s, _ := reader.ReadString('\n')
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			break
		}
		fmt.Printf("%d\n", fibonacci(n))
	}
}

func initdp() {
	dp[0] = big.NewInt(0)
	dp[1] = big.NewInt(1)
}

func fibonacci(n int) *big.Int {
	// base cases
	if n <= 1 || dp[n] != nil {
		return dp[n]
	}

	// cache solution
	dp[n-1] = fibonacci(n - 1)
	dp[n-2] = fibonacci(n - 2)
	dp[n] = &big.Int{}
	dp[n].Add(dp[n-1], dp[n-2])
	return dp[n]
}
