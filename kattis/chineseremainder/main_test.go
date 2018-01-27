package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	cases := []struct {
		a, n, b, m, x, k int
	}{
		{1, 2, 2, 3, 5, 6},
		{151, 783, 57, 278, 31471, 217674},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			x, k := chineseRemainder(tc.a, tc.n, tc.b, tc.m)
			assert.Equal(t, tc.x, x)
			assert.Equal(t, tc.k, k)
		})
	}
}
