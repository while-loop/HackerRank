package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		in           string
		correct, ttl int
	}{
		{"3 E right 10 A wrong 30 C wrong 50 B wrong 100 A wrong 200 A right 250 C wrong 300 D right -1", 3, 543},
		{"7 H right 15 B wrong 30 E wrong 35 E right 80 B wrong 80 B right 100 D wrong 100 C wrong 300 C right 300 D wrong -1", 3, 543},
	}
	for _, tc := range tcs {
		t.Run(tc.in, func(inner *testing.T) {
			correct, ttl := solve(bufio.NewScanner(strings.NewReader(tc.in)))
			assert.Equal(inner, tc.correct, correct)
			assert.Equal(inner, tc.ttl, ttl)
		})
	}
}
