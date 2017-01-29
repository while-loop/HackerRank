package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)

/**
  * https://open.kattis.com/contests/ze6qt6/problems/timebomb
  */
func main() {
	in.Split(bufio.ScanLines)

	nums := make([][5][3]byte, 8)
	for i := 0; i < 5; i++ {
		line := NextLine()
		for j, numIdx := 3, 0; j < len(line)+1; j, numIdx = j + 4, numIdx + 1 {
			nums[numIdx][i][0] = line[j - 3]
			nums[numIdx][i][1] = line[j - 2]
			nums[numIdx][i][2] = line[j - 1]
		}
	}

	finalNum := 0
	out := "BEER!!"
	for i := 0; i < len(nums); i++ {
		if nums[i][0][0] == 0 { // check if this num index has been init'd
			break
		}
		tmp := ascii2int(nums[i])
		if tmp < 0 {
			out = "BOOM!!"
			finalNum = -1
			break
		}
		finalNum = finalNum*10 + tmp
	}

	if finalNum%6 != 0 {
		out = "BOOM!!"
	}

	fmt.Println(out)
}

func ascii2int(ascii [5][3]byte) int {
	pos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(ascii) && (len(pos) > 1); i++ {
		pos = elim(i, ascii[i], pos)
	}

	ret := -1
	if len(pos) == 1 {
		ret = pos[0]
	}
	return ret
}

func elim(rowIdx int, vals [3]byte, possibles []int) []int {
	var localPos []int
	switch rowIdx {
	case 0:
		if exists(vals, "***") {
			localPos = []int{0, 2, 3, 5, 6, 7, 8, 9}
		} else if exists(vals, "* *") {
			localPos = []int{4}
		} else if exists(vals, "  *") {
			localPos = []int{1}
		}
		break
	case 1:
		if exists(vals, "* *") {
			localPos = []int{0, 4, 8, 9}
		} else if exists(vals, "  *") {
			localPos = []int{1, 2, 3, 7}
		} else if exists(vals, "*  ") {
			localPos = []int{5, 6}
		}
		break
	case 2:
		if exists(vals, "* *") {
			localPos = []int{0}
		} else if exists(vals, "  *") {
			localPos = []int{1, 7}
		} else if exists(vals, "***") {
			localPos = []int{2, 3, 4, 5, 6, 8, 9}
		}
		break
	case 3:
		if exists(vals, "* *") {
			localPos = []int{0, 6, 8}
		} else if exists(vals, "  *") {
			localPos = []int{1, 3, 4, 5, 7, 9}
		} else if exists(vals, "*  ") {
			localPos = []int{2}
		}
		break
	case 4:
		if exists(vals, "***") {
			localPos = []int{0, 2, 3, 5, 6, 8, 9}
		} else if exists(vals, "  *") {
			localPos = []int{1, 4, 7}
		}
		break
	}
	return intersect(localPos, possibles)
}

func exists(row [3]byte, vals string) bool {
	return row[0] == vals[0] && row[1] == vals[1] && row[2] == vals[2]
}

func intersect(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)

	n := len(a)
	m := len(b)

	min := n
	if m < min {
		min = m
	}

	arr := make([]int, 0)
	for i, j := 0, 0; i < n && j < m; {
		if a[i] < b[j] {
			i++
		} else if b[j] < a[i] {
			j++
		} else {
			arr = append(arr, a[i])
			i++
		}
	}
	return arr
}

// Helper Funcs
func NextInt() int {
	in.Scan()
	tmp, _ := strconv.Atoi(in.Text())
	return tmp
}

func NextLine() string {
	in.Scan()
	return in.Text()
}

func NextFloat64() float64 {
	in.Scan()
	tmp, _ := strconv.ParseFloat(in.Text(), 64)
	return tmp

}
