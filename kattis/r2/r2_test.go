package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		r1, r2, s int
	}{
		{11, 19, 15},
		{4, 2, 3},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.r2, solve(tc.r1, tc.s))
	}
}
