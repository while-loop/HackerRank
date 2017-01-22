package main

import (
	"bufio"
	"os"
	"fmt"
	"bytes"
	"strings"
)

/**
  * https://open.kattis.com/contests/e3hoii/problems/anewalphabet
  */
func main54() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	in.Scan(); input := in.Text()
	fmt.Println(ConvAlphabet(input))
}

func ConvAlphabet(s string) string {
	s = strings.ToLower(s)
	buf := bytes.Buffer{}

	for _, value := range s {
		transd := ALPHABET[value]
		if transd == "" {
			buf.WriteRune(value)
		} else {
			buf.WriteString(transd)
		}
	}

	return buf.String()
}

var ALPHABET = map[rune]string{
	'a': "@",
	'b': "8",
	'c': "(",
	'd': "|)",
	'e': "3",
	'f': "#",
	'g': "6",
	'h': "[-]",
	'i': "|",
	'j': "_|",
	'k': "|<",
	'l': "1",
	'm': "[]\\/[]",
	'n': "[]\\[]",
	'o': "0",
	'p': "|D",
	'q': "(,)",
	'r': "|Z",
	's': "$",
	't': "']['",
	'u': "|_|",
	'v': "\\/",
	'w': "\\/\\/",
	'x': "}{",
	'y': "`/",
	'z': "2",
}
