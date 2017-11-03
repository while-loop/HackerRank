package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUncoupled(t *testing.T) {
	in := "1, 2, 3, 5, 1, 2, 3"
	assert.Equal(t, 5, uncoupled(in), "Wrong Answer")
}

func TestUncoupled2(t *testing.T) {
	in := "1, 2, 3, 4, 5, 99, 1, 2, 3, 4, 5"
	assert.Equal(t, 99, uncoupled(in), "Wrong Answer")
}

func TestUncoupled3(t *testing.T) {
	in := "85, 1, 3, 2, 4, 5, 1, 2, 3, 4, 5"
	assert.Equal(t, 85, uncoupled(in), "Wrong Answer")
}

func TestCoupled(t *testing.T) {
	in := "2, 4, 4, 6, 8, 8, 2, 6"
	assert.Equal(t, 0, uncoupled(in), "Test not coupled")
}
