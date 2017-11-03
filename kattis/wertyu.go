package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/contests/ze6qt6/problems/wertyu
 */
func main() {
	in.Split(bufio.ScanLines)

	for {
		buf := bytes.Buffer{}
		line := NextLine()
		if line == "" {
			break
		}
		for i := 0; i < len(line); i++ {
			buf.WriteByte(xform(line[i]))
		}
		fmt.Println(buf.String())
	}
}

func xform(char byte) byte {
	if char >= '2' && char <= '9' {
		return char - 1
	} else {
		l := CHARS[char]
		if l == 0 {
			return char
		} else {
			return l
		}
	}
}

var CHARS = map[byte]byte{
	'W':  'Q',
	'E':  'W',
	'R':  'E',
	'T':  'R',
	'Y':  'T',
	'U':  'Y',
	'I':  'U',
	'O':  'I',
	'P':  'O',
	'[':  'P',
	']':  '[',
	'\\': ']',
	'S':  'A',
	'D':  'S',
	'F':  'D',
	'G':  'F',
	'H':  'G',
	'J':  'H',
	'K':  'J',
	'L':  'K',
	';':  'L',
	'\'': ';',
	'X':  'Z',
	'C':  'X',
	'V':  'C',
	'B':  'V',
	'N':  'B',
	'M':  'N',
	',':  'M',
	'.':  ',',
	'/':  '.',
	'-':  '0',
	'=':  '-',
	'1':  '`',
	'0':  '9',
}

func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp

}

func NextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp

}

func NextLine() string {
	in.Scan()
	return in.Text()
}
