/**
 * Set of all letters
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

/**

 */

func main() {
	reader := bufio.NewReader(os.Stdin)

	set := map[byte]bool{}
	for {
		t, err := reader.ReadByte()
		if t == 10 || err != nil {
			break
		}

		if t != ' ' {
			if 65 <= t && t <= 90 {
				t += 32
			}
			set[t] = true
		}
	}

	if len(set) == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
