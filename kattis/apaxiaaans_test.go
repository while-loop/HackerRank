package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExamplesCase0(t *testing.T) {

	in := []string{"bobapalaxiamethostenes", "bob", "boooooooooob", "boooob", "bbbooobbb", "booob", "boooooobapalaxxxxios"}
	exp := []string{"bobapalaxiamethostenes", "bob", "bob", "bob", "bob", "bob", "bobapalaxios"}

	for i := 0; i < len(in); i++ {
		assert.Equal(t, exp[i], Apaxias(in[i]), "FeelsUnluckyMan", i)
	}
}
