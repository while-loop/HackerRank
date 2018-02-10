package main

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		cur, pos, exp int
	}{
		{90, 359, -91},
		{315, 45, 90},
		{180, 270, 90},
		{45,270, -135},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, position(tc.cur, tc.pos))
		})
	}
}
