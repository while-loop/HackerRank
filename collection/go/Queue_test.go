package collection

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestQueue_EmptyQueueShouldBeZero(t *testing.T) {
	queue := Queue{}
	assert.Equal(t, 0, queue.Size(), "Queue size not 0")
	assert.True(t, queue.Empty(), "Queue is not empty")
}

func TestQueue_QueueSizeOne(t *testing.T) {
	queue := Queue{}
	queue.Push('(')
	assert.Equal(t, 1, queue.Size(), "Queue size not 1")
	assert.False(t, queue.Empty(), "Queue is empty")
}

func TestQueue_QueueSizeTwo(t *testing.T) {
	queue := Queue{}
	queue.Push('(')
	queue.Push(')')
	assert.Equal(t, 2, queue.Size(), "Queue size not 2")
	assert.False(t, queue.Empty(), "Queue is empty")
}

func TestQueue_ByteQueuePopFIFO(t *testing.T) {
	queue := Queue{}
	queue.Push(byte('*'))
	queue.Push(byte('^'))
	assert.Equal(t, 2, queue.Size(), "Queue size not 2")
	assert.False(t, queue.Empty(), "Queue is empty")

	val, err := queue.Pop()
	if err != nil {
		assert.Fail(t, "Failed to pop from queue")
	}
	assert.Equal(t, byte('*'), val, "Popped byte not FIFO")

	val, err = queue.Pop()
	if err != nil {
		assert.Fail(t, "Failed to pop from queue")
	}

	assert.Equal(t, byte('^'), val, "Popped byte not FIFO")
	_, err = queue.Pop()
	assert.NotNil(t, err, "Failed to get Empty Queue error")
}

func TestQueue_StringQueuePopFIFO(t *testing.T) {
	VAL1 := "testing"
	VAL2 := "0123456789"
	queue := Queue{}
	queue.Push(VAL1)
	queue.Push(VAL2)
	assert.Equal(t, 2, queue.Size(), "Queue size not 2")
	assert.False(t, queue.Empty(), "Queue is empty")

	val, err := queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")

	assert.Equal(t, VAL1, val.(string), "Popped byte not FIFO")

	val, err = queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")

	assert.Equal(t, VAL2, val.(string), "Popped byte not FIFO")

	_, err = queue.Pop()
	assert.NotNil(t, err, "Failed to get Empty Queue error")
}

func TestQueue_RealStringQueuePopFIFO(t *testing.T) {
	VAL1 := "testing"
	VAL2 := "0123456789"
	queue := StringQueue{}
	queue.Push(VAL1)
	queue.Push(VAL2)
	assert.Equal(t, 2, queue.Size(), "Queue size not 2")
	assert.False(t, queue.Empty(), "Queue is empty")

	val, err := queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")
	assert.Equal(t, VAL1, val, "Popped byte not FIFO")

	val, err = queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")
	assert.Equal(t, VAL2, val, "Popped byte not FIFO")

	_, err = queue.Pop()
	assert.NotNil(t, err, "Failed to get Empty Queue error")
}

func TestQueue_RealIntQueuePopFIFO(t *testing.T) {
	VAL1 := 123
	VAL2 := 8675309
	queue := IntQueue{}
	queue.Push(VAL1)
	queue.Push(VAL2)
	assert.Equal(t, 2, queue.Size(), "Queue size not 2")
	assert.False(t, queue.Empty(), "Queue is empty")

	val, err := queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")
	assert.Equal(t, VAL1, val, "Popped byte not FIFO")

	val, err = queue.Pop()
	assert.Nil(t, err, "Failed to pop from queue")
	assert.Equal(t, VAL2, val, "Popped byte not FIFO")

	_, err = queue.Pop()
	assert.NotNil(t, err, "Failed to get Empty Queue error")
}

func TestQueue_EmptyQueueReturnsErr(t *testing.T) {
	queue := Queue{}
	_, err := queue.Pop()
	assert.NotNil(t, err, "Failed to get Empty Queue error")
}

func TestQueue_Resize(t *testing.T) {
	queue := Queue{}
	for i := 0; i < 33; i++ {
		queue.Push(byte(i))
	}

	val, err := queue.Pop()
	assert.Nil(t, err, "Error when popping")
	assert.Equal(t, byte(0), val, "Popped wrong value")
	assert.Equal(t, 32, queue.Size(), "Wrong size when popped")
}

