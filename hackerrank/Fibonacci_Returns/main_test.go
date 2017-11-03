package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type Pair struct {
	a int
	b int64
}

func TestFib(t *testing.T) {
	initdp()
	vals := [14]Pair{{1, 1}, {8, 21}, {15, 610}, {22, 17711}, {29, 514229}, {36, 14930352}, {43, 433494437},
		{50, 12586269025}, {57, 365435296162}, {64, 10610209857723}, {71, 308061521170129}, {78, 8944394323791464},
		{85, 259695496911122585}, {92, 7540113804746346429}}
	for i := 0; i < len(vals); i++ {
		assert.Equal(t, big.NewInt(vals[i].b), fibonacci(vals[i].a), "Fibonacci not match", vals[i])
	}
}
