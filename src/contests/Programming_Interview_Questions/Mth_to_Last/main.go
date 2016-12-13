/**
  * Doubly LinkedList to solve problem. Keeps track of size in a variable and traverses from
  * either the head or tail of the list depending position of the given index.
  *
  */

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

/**
4
10 200 3 40000 5

200
 */
func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	index, _ := strconv.Atoi(strings.TrimSpace(s))

	s, _ = reader.ReadString('\n')
	ll := List(new(LinkedList))
	for _, element := range strings.Split(s, " ") {
		number, _ := strconv.Atoi(strings.TrimSpace(element))
		ll.Add(number)
	}

	val, err := ll.Get(ll.Size() - index)
	if err != nil {
		fmt.Println("NIL")
	} else {
		fmt.Println(val)
	}
}

type Node struct {
	prev  *Node;
	value int;
	next  *Node;
}

type List interface {
	Add(value int);
	Remove(index int) (int, error);
	Get(index int) (int, error);
	Empty() bool;
	Size() int;
}

type LinkedList struct {
	head *Node;
	tail *Node;
	size int;
}

func (this *LinkedList) Add(value int) {
	tmp := &Node{value:value, next:nil, prev:nil};
	if this.head == nil {
		this.head = tmp;
		this.tail = tmp;
		this.size = 0;
	} else {
		tmp.prev = this.tail; // assign previous in temp
		this.tail.next = tmp; // assign next in tail
		this.tail = tmp; // swap tail with temp
	}
	this.size++;
}

func (this *LinkedList) Remove(index int) (int, error) {
	tmp, err := this.getNode(index);
	if err != nil {
		return -1, err;
	} else {
		ret := tmp.value;
		if tmp.prev == nil { // head
			this.head = tmp.next;
		} else { // non-head
			tmp.prev.next = tmp.next;
		}

		if tmp.next == nil { // tail
			this.tail = tmp.prev;
		} else { // non-tail
			tmp.next.prev = tmp.prev;
		}

		tmp.next = nil;
		tmp.prev = nil;
		tmp = nil;
		this.size--;

		return ret, nil;
	}
}

/**
	Helper func to get a node at the given index.
	Does error checking for size.
 */
func (this *LinkedList) getNode(index int) (*Node, error) {
	if index < 0 || index >= this.size {
		return nil, errors.New("Index must be greater than 0 and less than LinkedList size");
	}

	// if we get this far, there exists at least one object in the list

	// Check the size of the list and traverse from tail or head. depending on which is closer to the given index
	tmp := &Node{}
	useHead := true
	if (this.size - index) > (this.size / 2.0) {
		// index is closer to head
		tmp = this.head;
	} else {
		// index is closer to tail
		tmp = this.tail;
		index = this.size - index - 1;
		useHead = false
	}

	for tmp != nil && 0 < index {
		if useHead {
			tmp = tmp.next;
		} else {
			tmp = tmp.prev;
		}
		index--;
	}

	if tmp == nil {
		return nil, errors.New("Uh oh, less items in list than list.size dictates..");
	} else {
		return tmp, nil;
	}
}

func (this *LinkedList) Get(index int) (int, error) {
	tmp, err := this.getNode(index);
	if err != nil {
		return -1, err;
	}

	return tmp.value, nil;

}

func (this *LinkedList)  Print() {
	tmp := this.head;
	for tmp != nil {
		fmt.Printf("%d, ", tmp.value)
		tmp = tmp.next;
	}
	fmt.Println()
}

func (this *LinkedList) Empty() bool {
	return this.size == 0;
}

func (this *LinkedList) Size() int {
	return this.size;
}