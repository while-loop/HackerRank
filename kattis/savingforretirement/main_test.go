package main

import (
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		bAge, bRetire, bSave, aAge, aSave, aRetire int
	}{
		{20, 25, 5, 20, 10, 23},
		{20, 28, 5, 30, 9, 35},
		{20, 28, 5, 20, 4, 31},
		{20, 28, 5, 20, 100, 21},
		{20, 100, 100, 20, 1, 8021},
	}
	for idx, tc := range tcs {
		t.Run(strconv.Itoa(idx), func(inner *testing.T) {
			assert.Equal(inner, tc.aRetire, breakEven(tc.bAge, tc.bRetire, tc.bSave, tc.aAge, tc.aSave))
		})
	}
}
