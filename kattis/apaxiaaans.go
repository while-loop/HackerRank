package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/apaxiaaans
 */
func main85() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	word := in.Text()
	fmt.Println(Apaxias(word))
}

func Apaxias(s string) string {
	buf := bytes.Buffer{}

	var prevL rune
	for _, letter := range s {
		if letter != prevL {
			buf.WriteRune(letter)
		}
		prevL = letter
	}
	return buf.String()
}
