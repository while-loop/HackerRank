package main

import (
	"fmt"
	"reflect"
)

func main() {
	exp := []int{1, 2}
	act := twoSum([]int{2, 7, 11, 15}, 9)
	if reflect.DeepEqual(exp, act) {
		fmt.Printf("%v == %v\n", exp, act)
	} else {
		fmt.Printf("%v != %v\n", exp, act)
	}
}

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1

	for low < high {
		result := numbers[low] + numbers[high]
		if result == target {
			return []int{low + 1, high + 1}
		} else if result < target {
			low++
		} else {
			high--
		}
	}

	return nil
}
