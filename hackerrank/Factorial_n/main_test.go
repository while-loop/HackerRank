package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/big"
)

type Pair struct {
	a int
	b int64
}

func TestCases(t *testing.T) {
	vals := []Pair{{3, 6}, {0, 1}, {7, 5040}, {10, 3628800}, {20, 2432902008176640000}}

	for i := 0; i < len(vals); i++ {
		assert.Equal(t, big.NewInt(vals[i].b), factorial(vals[i].a), "Factorial not match", vals[i])
	}
}