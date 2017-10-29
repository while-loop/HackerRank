package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		statuses, days int
	}{
		{1, 1},
		{5, 4},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.days, solve(tc.statuses))
	}
}
