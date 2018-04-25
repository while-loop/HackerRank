package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	cases := []struct {
		in  []int
		exp [][]int
	}{
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, tc := range cases {
		act := threeSum(tc.in)
		if reflect.DeepEqual(tc.exp, act) {
			fmt.Printf("pass\n%v == %v\n", tc.exp, act)
		} else {
			fmt.Printf("fail\n%v != %v\n", tc.exp, act)
		}
		fmt.Println()
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)

	for idx, val := range nums {
		if idx != 0 && !(idx > 0 && nums[idx] != nums[idx-1]) {
			continue
		}

		target := 0 - val
		low, high := idx+1, len(nums)-1
		for low < high {
			if nums[low]+nums[high] == target {
				results = append(results, []int{nums[idx], nums[low], nums[high]})

				// move low pointer to non duplicate low
				for low < high && nums[low] == nums[low+1] {
					low++
				}

				// move high pointer to non duplicate low
				for low < high && nums[high] == nums[high-1] {
					high--
				}

				low++
				high--
			} else if nums[low]+nums[high] < target {
				low++
			} else {
				high--
			}
		}
	}
	return results
}
