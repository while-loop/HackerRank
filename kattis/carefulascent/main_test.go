package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		x, y    float64
		shields []Shield
		exp     float64
	}{
		{100, 140, []Shield{
			{40, 90, 0.2000000000},
		}, 1.0},
		{100, 100, []Shield{
			{0, 20, 2.0000000000},
			{50, 100, 0.1000000000},
			{20, 50, 0.2000000000},
		}, 1.96078431373},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			act, err := strconv.ParseFloat(fmt.Sprintf("%.11f", horizontalVelocity(tc.x, tc.y, tc.shields)), 64)
			assert.NoError(t, err)
			assert.Equal(inner, tc.exp, act)
		})
	}
}
