package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		arr             []int
		multiplications int
	}{
		{[]int{1, 2, 3, 4, 3}, 30},
		{[]int{10, 30, 5, 60}, 4500},
		{[]int{40, 20, 30, 10, 30}, 26000},
		{[]int{10, 20, 30, 40, 30}, 30000},
		{[]int{10, 20, 30}, 6000},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			assert.Equal(t, tc.multiplications, MatrixChainOrder(tc.arr))
		})
	}
}

func TestDP(t *testing.T) {
	arr := []int{10, 20, 30, 40, 30, 10, 20, 30, 40, 30, 40, 30, 40, 40, 30, 10, 20, 30, 40, 30}
	calcs := 161000
	assert.Equal(t, calcs, MatrixChainOrder(arr))
}
