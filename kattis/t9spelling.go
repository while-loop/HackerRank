package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/t9spelling
 */
func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	in.Scan()
	N, _ := strconv.Atoi(in.Text())
	for i := 1; i <= N; i++ {
		in.Scan()
		fmt.Printf("Case #%d: %s\n", i, Word2Nums(in.Text()))
	}
}

func Word2Nums(s string) string {
	s = strings.ToLower(s)
	buf := bytes.Buffer{}

	var lastDigit byte
	for _, letter := range s {

		pos := letter - 'a' + 1 // zero base the letter
		if letter == ' ' {
			pos = 0
		}

		num := NUMBER[pos]
		counts := REPEATS[pos]

		if lastDigit == num {
			buf.WriteRune(' ')
		}
		lastDigit = num

		for i := 0; i < counts; i++ {
			buf.WriteByte(num)
		}
	}
	return buf.String()
}

var (
	REPEATS = []int{1, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 4, 1, 2, 3, 1, 2, 3, 4}
	NUMBER  = "022233344455566677778889999"
)
