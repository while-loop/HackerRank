package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	vals := [17]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 18, 20, 22, 24, 26, 28, 30, 35}
	expected := [17]int{1, 2, 5, 13, 34, 89, 233, 610, 1597, 2584, 6765, 17711, 46368, 121393, 317811, 832040, 9227465}
	for i := 0; i < len(vals); i++ {
		assert.Equal(t, expected[i], fibonacci(vals[i]), "Fibonacci not match", vals[i])
	}
}
