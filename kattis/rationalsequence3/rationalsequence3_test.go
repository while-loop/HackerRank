package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	tcs := []struct {
		name string
		n    int
		p, q uint
	}{
		{"1", 1, 1, 1},
		{"4", 4, 1, 3},
		{"11", 11, 5, 2},
		{"1431655765", 1431655765, 2178309, 1346269},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(inner *testing.T) {
			p, q := Solve(tc.n)
			str:=fmt.Sprintf("%v", tc)
			assert.Equal(inner, int(tc.p), int(p), str)
			assert.Equal(inner, int(tc.q), int(q), str)
		})
	}
}
