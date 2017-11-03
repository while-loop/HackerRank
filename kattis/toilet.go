package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * https://open.kattis.com/contests/e3hoii/problems/toilet
 */
func main5454() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	in.Scan()
	word := in.Text()

	changes := make([]int, 3)
	lastPos := make([]byte, 3)
	lastPos[0] = word[0]
	lastPos[1] = word[0]
	lastPos[2] = word[0]
	for i := 1; i < len(word); i++ {
		curPos := word[i]

		// pol 1
		if lastPos[0] == 'D' {
			changes[0] += 1
		} else {
			if curPos == 'D' {
				changes[0] += 2
			}
		}
		lastPos[0] = 'U'

		// pol 2
		if lastPos[1] == 'D' {
			if curPos == 'U' {
				changes[1] += 2
			}
		} else {
			changes[1] += 1
		}
		lastPos[1] = 'D'

		// pol 3
		if lastPos[2] != curPos {
			changes[2] += 1
		}
		lastPos[2] = curPos
	}

	fmt.Printf("%d\n%d\n%d\n", changes[0], changes[1], changes[2])
}
