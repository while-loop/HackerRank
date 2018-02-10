package main

import (
	"testing"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		mboxs, fworks, exp int
	}{
		{1, 10, 55},
		{1, 100, 5050},
		{3, 73, 382},
		{5, 100, 495},
	}

	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.exp, totalFireworks(tc.mboxs, tc.fworks))
		})
	}
}
