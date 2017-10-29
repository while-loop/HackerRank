package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	changes := []int{3, 2, -1, 1, 2}
	max, denied := 4, 2

	sem := newSem(max)
	for _, change := range changes {
		sem.change(change)
	}

	assert.Equal(t, denied, sem.denies)
}
