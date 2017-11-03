package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestStack_ByteStackPopLIFO(t *testing.T) {
	stack := Stack{}
	stack.Push(byte('*'))
	stack.Push(byte('^'))
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

func TestStack_StringStackPopLIFO(t *testing.T) {
	VAL1 := "testing"
	VAL2 := "0123456789"
	stack := Stack{}
	stack.Push(VAL1)
	stack.Push(VAL2)
	assert.Equal(t, 2, stack.Size(), "Stack size not 2")
	assert.False(t, stack.Empty(), "Stack is empty")

	val, err := stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")

	assert.Equal(t, VAL2, val.(string), "Popped byte not LIFO")

	val, err = stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")

	assert.Equal(t, VAL1, val.(string), "Popped byte not LIFO")

	_, err = stack.Pop()
	assert.NotNil(t, err, "Failed to get Empty Stack error")
}

func TestStack_RealStringStackPopLIFO(t *testing.T) {
	VAL1 := "testing"
	VAL2 := "0123456789"
	stack := StringStack{}
	stack.Push(VAL1)
	stack.Push(VAL2)
	assert.Equal(t, 2, stack.Size(), "Stack size not 2")
	assert.False(t, stack.Empty(), "Stack is empty")

	val, err := stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")
	assert.Equal(t, VAL2, val, "Popped byte not LIFO")

	val, err = stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")
	assert.Equal(t, VAL1, val, "Popped byte not LIFO")

	_, err = stack.Pop()
	assert.NotNil(t, err, "Failed to get Empty Stack error")
}

func TestStack_RealIntStackPopLIFO(t *testing.T) {
	VAL1 := 123
	VAL2 := 8675309
	stack := IntStack{}
	stack.Push(VAL1)
	stack.Push(VAL2)
	assert.Equal(t, 2, stack.Size(), "Stack size not 2")
	assert.False(t, stack.Empty(), "Stack is empty")

	val, err := stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")
	assert.Equal(t, VAL2, val, "Popped byte not LIFO")

	val, err = stack.Pop()
	assert.Nil(t, err, "Failed to pop from stack")
	assert.Equal(t, VAL1, val, "Popped byte not LIFO")

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
}

func TestStack_Peek(t *testing.T) {
	stack := Stack{}
	for i := 0; i < 33; i++ {
		stack.Push(byte(i))
	}

	peeked, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, byte(32), peeked)

	peeked, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, byte(32), peeked)

	popped, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, byte(32), popped)

	peeked, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, byte(31), peeked)

	stack.Push(byte(55))

	peeked, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, byte(55), peeked)
	assert.Equal(t, 33, stack.Size(), "Wrong size when popped")

}
