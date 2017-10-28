package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		in, out string
	}{
		{"Knuth-Morris-Pratt", "KMP"},
		{"Mirko-Slavko", "MS"},
		{"Pasko-Patak", "PP"},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(inner *testing.T) {
			assert.Equal(inner, tc.out, solve(tc.in))
		})
	}
}
