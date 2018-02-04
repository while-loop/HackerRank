package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		mos, pup, larv, eggs, rform, sAdult, n int
		exp                                    int
	}{
		{10, 20, 40, 4, 2, 2, 10, 10},
		{144, 55, 8, 0, 1, 9, 4, 0},
		{10, 10, 10, 2, 3, 2, 6, 1},
		{10, 20, 40, 86, 9, 9, 999, 10},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, solve(tc.mos, tc.pup, tc.larv, tc.eggs, tc.rform, tc.sAdult, tc.n))
		})
	}
}
