package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		boxes  []int
		h1, h2 int
		exp    []int
	}{
		{[]int{12, 8, 2, 4, 10, 3}, 25, 14, []int{12, 10, 3, 8, 4, 2}},
		{[]int{12, 8, 2, 4, 10, 3}, 14, 25, []int{8, 4, 2, 12, 10, 3}},
		{[]int{12, 17, 36, 37, 51, 63}, 92, 124, []int{63, 17, 12, 51, 37, 36}},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, boxConfig(tc.boxes, tc.h1, tc.h2))
		})
	}
}
