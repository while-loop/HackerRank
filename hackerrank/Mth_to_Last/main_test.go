package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	ll := List(new(LinkedList))
	count := 20
	for i := 0; i < count; i++ {
		ll.Add(i)
	}

	for i := 0; i < count; i++ {
		num, err := ll.Get(i)
		if err != nil {
			assert.Error(t, err)
		}

		assert.Equal(t, i, num, "LinkedList Get() Out of order")
	}
}

func TestRemove(t *testing.T) {
	ll := List(new(LinkedList))
	count := 20
	for i := 0; i < count; i++ {
		ll.Add(i)
	}

	ll.Remove(0)
	ll.Remove(ll.Size() - 1)
	ll.Remove(ll.Size() / 2)
	expected := [17]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18}
	assert.Equal(t, 17, ll.Size(), "LinkedList did not remove items")
	for i := 0; i < ll.Size(); i++ {
		num, err := ll.Get(i)
		if err != nil {
			assert.Error(t, err, "Error getting item from LinkedList")
		}

		assert.Equal(t, expected[i], num, "LinkedList Get() Out of order")
	}
}

func TestLinkedList_Empty(t *testing.T) {
	ll := List(new(LinkedList))
	assert.True(t, ll.Empty(), "LinkedList not empty")
}

func TestLinkedList_Not_Empty(t *testing.T) {
	ll := List(new(LinkedList))
	ll.Add(4)
	ll.Add(5)
	ll.Add(6)
	ll.Remove(1)
	assert.False(t, ll.Empty(), "LinkedList is empty")
}

func TestLinkedList_Size(t *testing.T) {
	ll := List(new(LinkedList))
	count := 6
	for i := 0; i < count; i++ {
		ll.Add(i)
	}

	ll.Remove(0)
	count--
	ll.Remove(ll.Size() - 1)
	count--

	assert.Equal(t, count, ll.Size(), "LinkedList size not equal")
}

func TestLinkedList_Remove_Empty(t *testing.T) {
	ll := List(new(LinkedList))
	_, err := ll.Remove(0)
	if err == nil {
		assert.Fail(t, "No error thrown")
	}

	assert.Equal(t, 0, ll.Size(), "LinkedList not empty")
}

func TestLinkedList_ProblemTest(t *testing.T) {
	index := 4
	vals := []int{10, 200, 3, 40000, 5}
	ll := List(new(LinkedList))
	for _, num := range vals {
		ll.Add(num)
	}

	val, err := ll.Get(ll.Size() - index)
	if err != nil {
		assert.Fail(t, "Unable to Get index: ", index)
	} else {
		assert.Equal(t, 200, val, "Wrong value for test case: ", val)
	}
}
