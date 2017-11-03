package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Pair struct {
	a int
	b int64
}

func TestMakeChange0(t *testing.T) {
	N := 4
	M := 3
	Cs := []int{1, 2, 3}
	exp := 4

	assert.Equal(t, exp, MakeChange(N, M, Cs), "Incorrect change amount", N, M, Cs)
}

func TestMakeChange1(t *testing.T) {
	N := 10
	M := 4
	Cs := []int{2, 5, 3, 6}
	exp := 5

	assert.Equal(t, exp, MakeChange(N, M, Cs), "Incorrect change amount", N, M, Cs)
}
