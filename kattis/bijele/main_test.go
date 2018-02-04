package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		given, exp []int
	}{
		{[]int{0, 1, 2, 2, 2, 7}, []int{1, 0, 0, 0, 0, 1}},
		{[]int{2, 1, 2, 1, 2, 1}, []int{-1, 0, 0, 1, 0, 7}},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, validate(tc.given))
		})
	}
}
