package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		a, b int
		exp  []int
	}{
		{6, 6, []int{7}},
		{6, 4, []int{5, 6, 7}},
		{12, 20, []int{13, 14, 15, 16, 17, 18, 19, 20, 21}},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, outcomes(tc.a, tc.b))
		})
	}
}
