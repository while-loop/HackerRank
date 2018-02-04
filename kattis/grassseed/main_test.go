package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		c   float64
		as  []rect
		exp float64
	}{
		{2, []rect{{3, 2}, {5, 4}, {6, 5}}, 112.0000000},
		{0.75, []rect{{2, 3.333}, {4.567, 3.41}}, 16.6796025},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, totalCost(tc.c, tc.as))
		})
	}
}
