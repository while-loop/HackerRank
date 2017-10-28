package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/autori
 */
func main() {
	in.Split(bufio.ScanWords)
	fmt.Println(solve(nextLine()))
}

func solve(name string) string {
	buf := bytes.NewBufferString("")
	for _, val := range strings.Split(name, "-") {
		buf.WriteByte(val[0])
	}

	return buf.String()
}

func nextLine() string {
	in.Scan()
	return in.Text()
}
