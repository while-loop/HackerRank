package main

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		low, high, exp int64
	}{
		{0, 10, 46},
		{28, 31, 28},
		{1234, 56789, 1128600},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, sum(tc.low, tc.high))
		})
	}
}

func TestSum(t *testing.T) {
	tcs := []struct {
		n, exp int64
	}{
		{84001, 13},
		{10000000000, 1},

	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, digitSum(tc.n))
		})
	}
}
