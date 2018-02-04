package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	totalBoxes = 6
)

// https://open.kattis.com/contests/dasomz/problems/towering
func main() {
	in.Split(bufio.ScanWords)

	boxes := make([]int, totalBoxes)
	for i := 0; i < totalBoxes; i++ {
		boxes[i] = nextInt()
	}

	for _, v := range boxConfig(boxes, nextInt(), nextInt()) {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// solve any box set using the alg to find a pair of numbers that equal a sum minus a random number
// then use the numbers not included in that set to equal the other box set
func boxConfig(boxes []int, h1, h2 int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(boxes)))

	victim := boxes[0]
	n1, n2 := findPair(boxes[1:], h1-victim)
	ret := make([]int, totalBoxes)
	if n1 == -1 { // wasn't found in this pair
		n1, n2 = findPair(boxes[1:], h2-victim)
		ret[3], ret[4], ret[5] = victim, n1, n2
		pos := 0
		for i := 0; i < totalBoxes; i++ {
			switch boxes[i] {
			case victim, n1, n2:
				break
			default:
				ret[pos] = boxes[i]
				pos++
			}
		}
	} else {
		ret[0], ret[1], ret[2] = victim, n1, n2
		pos := 3
		for i := 0; i < totalBoxes; i++ {
			switch boxes[i] {
			case victim, n1, n2:
				break
			default:
				ret[pos] = boxes[i]
				pos++
			}
		}
	}

	return ret
}

// findPair gets two numbers that equal the given sum.
// nums MUST be sorted in desc order.
// returns higher number first
func findPair(nums []int, sum int) (int, int) {
	start, end := 0, len(nums)-1
	for start < end {
		s := nums[start] + nums[end]
		if s == sum {
			return nums[start], nums[end] // return higher first
		} else if s > sum {
			start++
		} else {
			end--
		}
	}
	return -1, -1
}

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	tmp, err := strconv.Atoi(in.Text())
	if err != nil {
		panic(err)
	}
	return tmp
}
