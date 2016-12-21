package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"bytes"
	"math"
)

/**
  * Wishing you Godspeed
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/speed-7/
  */
func main12() {
	fmt.Print(SolveSpeed(bufio.NewScanner(os.Stdin)))
}

func SolveSpeed(scanner *bufio.Scanner) string {
	scanner.Split(bufio.ScanWords)
	var buffer bytes.Buffer

	scanner.Scan(); tests, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < tests; i++ {
		scanner.Scan(); cars, _ := strconv.Atoi(scanner.Text())
		max := math.MaxInt32
		count := 0
		for i := 0; i < cars; i++ {
			scanner.Scan(); cur, _ := strconv.Atoi(scanner.Text())

			if cur < max {
				max = cur
				count++
			}
		}
		buffer.WriteString(fmt.Sprintf("%d\n", count))
	}
	return buffer.String()
}