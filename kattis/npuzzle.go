package main

import (
	"bufio"
	"os"
	"fmt"
)

var in = bufio.NewScanner(os.Stdin)
/**
  * https://open.kattis.com/contests/ze6qt6/problems/npuzzle
  */
func main() {
	in.Split(bufio.ScanBytes)

	pts := 0
	for i := 0; i < 16; {
		val := NextByte()
		if val >= 'A' && val <= 'Z' {
			offset := int(val - 'A')

			xOff := abs((offset % 4) - (i % 4))
			yOff := abs((offset / 4) - (i / 4))
			pts += xOff + yOff
			i++
		} else if val == 0 {
			break
		} else if val == '.' {
			i++
		}
	}
	fmt.Println(pts)

}
func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

func NextByte() byte {
	if !in.Scan() {
		return 0
	}
	return in.Bytes()[0]
}
