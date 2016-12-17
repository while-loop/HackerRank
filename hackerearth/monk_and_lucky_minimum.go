package main

import (
	"bufio"
	"os"
	"strconv"
	"math"
	"fmt"
	"bytes"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-lucky-minimum-3/
  */
func asfd() {
	fmt.Print(SolveMonk(bufio.NewScanner(os.Stdin)))
}

func SolveMonk(scanner *bufio.Scanner) string {
	var buffer bytes.Buffer

	scanner.Split(bufio.ScanWords)
	scanner.Scan(); trials, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < trials; i++ {
		scanner.Scan(); nums, _ := strconv.Atoi(scanner.Text())
		var tmp int
		min, freq := math.MaxInt32, 0
		for j := 0; j < nums; j++ {
			scanner.Scan(); tmp, _ = strconv.Atoi(scanner.Text())
			if tmp == min {
				freq++
			} else if tmp < min {
				min = tmp
				freq = 1
			}
		}
		if freq % 2 == 0 {
			buffer.WriteString("Unlucky\n")
		} else {
			buffer.WriteString("Lucky\n")
		}
	}

	return buffer.String()
}