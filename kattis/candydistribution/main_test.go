package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	cases := []struct {
		k, c, exp int
	}{
		{10, 5, -1},
		{10, 7, 3},
		{1337, 23, 872},
		{123454321, 42, 14696943},
		{999999937, 142857133, 166666655},
	}

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.exp, candyDistribtion(tc.k, tc.c))
		})
	}
}
