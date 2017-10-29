package main

import (
	"fmt"
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
		fmt.Println("sd")
		assert.Equal(t, tc.days, solve(tc.statuses))
	}
}
