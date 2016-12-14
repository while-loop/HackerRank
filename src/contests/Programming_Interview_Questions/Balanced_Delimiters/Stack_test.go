package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStack_EmptyStackShouldBeZero(t *testing.T) {
	stack := Stack{}
	assert.Equal(t, 0, stack.Size(), "Stack size not 0")
	assert.True(t, stack.Empty(), "Stack is not empty")
}

func TestStack_StackSizeOne(t *testing.T) {
	stack := Stack{}
	stack.Push('(')
	assert.Equal(t, 1, stack.Size(), "Stack size not 1")
	assert.False(t, stack.Empty(), "Stack is empty")
}

func TestStack_StackSizeTwo(t *testing.T) {
	stack := Stack{}
	stack.Push('(')
	stack.Push(')')
	assert.Equal(t, 2, stack.Size(), "Stack size not 2")
	assert.False(t, stack.Empty(), "Stack is empty")
}

func TestStack_StackPopLIFO(t *testing.T) {
	stack := Stack{}
	stack.Push('*')
	stack.Push('^')
	assert.Equal(t, 2, stack.Size(), "Stack size not 2")
	assert.False(t, stack.Empty(), "Stack is empty")

	val, err := stack.Pop()
	if err != nil {
		assert.Fail(t, "Failed to pop from stack")
	}

	assert.Equal(t, byte('^'), val, "Popped byte not LIFO")

	val, err = stack.Pop()
	if err != nil {
		assert.Fail(t, "Failed to pop from stack")
	}

	assert.Equal(t, byte('*'), val, "Popped byte not LIFO")

	_, err = stack.Pop()
	assert.NotNil(t, err, "Failed to get Empty Stack error")
}

func TestStack_EmptyStackReturnsErr(t *testing.T) {
	stack := Stack{}
	_, err := stack.Pop()
	assert.NotNil(t, err, "Failed to get Empty Stack error")
}

func TestStack_Resize(t *testing.T) {
	stack := Stack{}
	for i := 0; i < 33; i++ {
		stack.Push(byte(i))
	}

	val, err := stack.Pop()
	assert.Nil(t, err, "Error when popping")
	assert.Equal(t, byte(32), val, "Popped wrong value")
	assert.Equal(t, 32, stack.Size(), "Wrong size when popped")
	assert.Equal(t, 64, len(stack.elements), "Wrong resize length")
}

