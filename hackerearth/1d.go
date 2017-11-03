package main

import "fmt"

/*
https://www.hackerearth.com/practice/data-structures/arrays/1-d/tutorial/
*/
func gsg5() {
	var n int
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(nums[i])
	}
}
