package main

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{} // [num] = idx

	for idx, v := range nums {
		needed := target - v
		if counterIdx, exists := numMap[needed]; exists && counterIdx != idx {
			return []int{counterIdx, idx}
		}

		numMap[v] = idx
	}

	return nil
}
