package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

/**
 * https://open.kattis.com/problems/apples
 */
func main() {
	in.Split(bufio.ScanWords)
	rows, cols := nextInt(), nextInt()

	arr := make([][]byte, rows)
	for r := 0; r < rows; r++ {
		arr[r] = make([]byte, cols)
		str := next()
		for c := 0; c < cols; c++ {
			arr[r][c] = str[c]
		}
	}
	solve(arr)
	printArr(arr)
}

func solve(in [][]byte) {
	for c := len(in[0]) - 1; c >= 0; c-- {
		dropApples(in, c)
	}
}

func dropApples(in [][]byte, col int) {
	prevRow := len(in) - 1

	// count how many apples seen until we've hit the top of the
	// matrix or `#` is found
	for prevRow >= 0 {
		applesFound := 0
		blockRow := prevRow

		for ; blockRow >= 0 && in[blockRow][col] != '#'; blockRow-- {
			if in[blockRow][col] == 'a' {
				applesFound++
				in[blockRow][col] = '.'
			}
		}

		// drop apples from `blockRow` up
		for r := 0; r < applesFound; r++ {
			in[prevRow-r][col] = 'a'
		}

		prevRow = blockRow - 1
	}
}

func printArr(in [][]byte) {
	for r := 0; r < len(in); r++ {
		for c := 0; c < len(in[r]); c++ {
			fmt.Printf("%c", in[r][c])
		}
		fmt.Println()
	}
}

func next() string {
	in.Scan()
	return in.Text()
}

func nextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}
