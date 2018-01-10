package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		n        uint
		expected int
	}{
		{17, 35},
		{7, 12},
		{45, 119},
		{127, 448},
	}

	for idx, tc := range tcs {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			assert.Equal(t, tc.expected, CountSetBits(tc.n))
		})
	}
}
