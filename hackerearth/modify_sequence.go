package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/modify-sequence/
 */
func rght() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < n-1; i++ {
		if arr[i] < 0 {
			fmt.Println("NO")
			os.Exit(0)
		}

		arr[i+1] -= arr[i]
		arr[i] = 0
	}

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			fmt.Println("NO")
			os.Exit(0)
		}
	}

	fmt.Println("YES")
}
