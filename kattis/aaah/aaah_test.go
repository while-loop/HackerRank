package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		given, required, shouldGo string
	}{
		{"aaah", "aaaaah", "no"},
		{"aaah", "ah", "go"},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.shouldGo, solve(tc.given, tc.required))
	}
}
