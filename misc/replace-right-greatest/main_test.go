package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		in, out []int
	}{
		{[]int{16, 17, 4, 3, 5, 2}, []int{17, 5, 5, 5, 2, -1}},
		{[]int{10, 30, 5, 60}, []int{60, 60, 60, -1}},
		{[]int{40, 20, 30, 10, 30}, []int{30, 30, 30, 30, -1}},
		{[]int{10, 20, 30, 40, 30}, []int{40, 40, 40, 30, -1}},
		{[]int{16, 17, 4, 3, 5, 2}, []int{17, 5, 5, 5, 2, -1}},
		{[]int{2, 3, 1, 9}, []int{9, 9, 9, -1}},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			assert.Equal(t, tc.out, ReplaceRightGreatest(tc.in))
		})
	}
}
