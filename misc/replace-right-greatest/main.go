package main

import (
	"fmt"
)

// https://www.geeksforgeeks.org/replace-every-element-with-the-greatest-on-right-side/
func main() {
	arr := []int{}
	fmt.Println(ReplaceRightGreatest(arr))
}

// ReplaceRightGreatest is an in-place func to replace every element with the next
// greatest element (greatest element on the right side) in the array
func ReplaceRightGreatest(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[len(arr)-1]
	currentMax := max

	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > currentMax {
			currentMax = arr[i]
		}

		arr[i] = max
		max = currentMax
	}

	return arr
}
