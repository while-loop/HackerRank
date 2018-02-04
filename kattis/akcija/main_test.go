package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		books []int
		cost  int64
	}{
		{[]int{3, 2, 3, 2}, 8},
		{[]int{6, 4, 5, 5, 5, 5}, 21},
		{[]int{}, 0},
		{[]int{100}, 100},
		{[]int{100000, 100000, 100000, 100000}, 300000},
		{[]int{100000, 100000, 100000, 100000}, 300000},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.cost, totalCost(tc.books))
		})
	}

	books := make([]int, 0)
	its := 1432
	bulk := 15
	cost := 100000
	for i := 0; i < its; i++ {
		books = append(books, []int{
			cost, cost, cost,
			cost, cost, cost,
			cost, cost, cost,
			cost, cost, cost,
			cost, cost, cost,
		}...)

	}
	ttl := int64(cost*bulk*its) - int64(cost*(bulk/3)*its)
	assert.Equal(t, ttl, totalCost(books))
}
