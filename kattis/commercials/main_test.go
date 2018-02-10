package main

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		cost int
		breaks []int
		exp int
	}{
		{20, []int{18, 35, 6, 80, 15, 21}, 61},
		{20, []int{40, 40, 1, 1, 1, 5}, 40},
		{20, []int{1, 100, 0, 80, 1, 5}, 120},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, maxProfit(tc.cost, tc.breaks))
		})
	}
}
