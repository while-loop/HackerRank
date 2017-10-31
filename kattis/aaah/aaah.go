package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/aaah
 */
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(solve(next(), next()))
}

// given: Jon's ah
// required: doctor's ah
func solve(given, required string) string {
	requiredIdx := strings.Index(required, "h")
	givenIdx := strings.Index(given, "h")

	if givenIdx >= requiredIdx {
		return "go"
	}

	return "no"
}

func next() string {
	in.Scan()
	return in.Text()
}
