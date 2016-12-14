package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	n := 1
	exp := []string{"1"}
	arr := FiZzBuZz(n)
	assert.Equal(t, exp, arr, "Incorrect fizzbuzz output", n)
}

func TestCase2(t *testing.T) {
	n := 15
	exp := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	arr := FiZzBuZz(n)
	assert.Equal(t, exp, arr, "Incorrect fizzbuzz output", n)
}

func TestCase3(t *testing.T) {
	n := 22
	exp := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
	"16", "17", "Fizz", "19", "Buzz", "Fizz", "22"}
	arr := FiZzBuZz(n)
	assert.Equal(t, exp, arr, "Incorrect fizzbuzz output", n)
}
