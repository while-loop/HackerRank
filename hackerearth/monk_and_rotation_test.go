package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRotationCase0(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	shift := 2
	exp := "4 5 1 2 3 "

	assert.Equal(t, exp, SolveRotation(arr, shift))
}

func TestRotationCase1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	shift := 10
	exp := "1 2 3 4 5 "

	assert.Equal(t, exp, SolveRotation(arr, shift))
}

func TestRotationCase2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	shift := 3
	exp := "3 4 5 1 2 "

	assert.Equal(t, exp, SolveRotation(arr, shift))
}

func TestRotationCase3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	shift := 5
	exp := "1 2 3 4 5 "

	assert.Equal(t, exp, SolveRotation(arr, shift))
}