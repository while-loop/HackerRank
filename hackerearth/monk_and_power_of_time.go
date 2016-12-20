package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"errors"
)

/**
  * https://www.hackerearth.com/practice/data-structures/arrays/1-d/practice-problems/algorithm/monk-and-power-of-time/
  */
func asd() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan(); n, _ := strconv.Atoi(scanner.Text())
	idealqueue := make([]int, n)
	givenqueue := IntQueue{}

	for i := 0; i < n; i++ {
		scanner.Scan(); tmp, _ := strconv.Atoi(scanner.Text())
		givenqueue.Push(tmp)
	}

	for i := 0; i < n; i++ {
		scanner.Scan(); idealqueue[i], _ = strconv.Atoi(scanner.Text())
	}

	ttl := 0
	for i := 0; i < n; i++ {
		for {
			val, _ := givenqueue.Pop()
			if val == idealqueue[i] {
				break
			}
			ttl++
			givenqueue.Push(val)
		}
		ttl++
	}
	fmt.Println(ttl)
}


/**
	Kinda generic Queue
 */
type Queue struct {
	elements []interface{}
}

func (this *Queue) Empty() bool {
	return len(this.elements) == 0
}

func (this *Queue) Peek() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Queue")
	} else {
		return this.elements[0], nil
	}
}

func (this *Queue) Pop() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Queue")
	} else {
		res := this.elements[0]
		this.elements[0] = nil
		this.elements = this.elements[1:]
		return res, nil
	}
}

func (this *Queue) Push(val interface{}) {
	this.elements = append(this.elements, val)
}

func (this *Queue) Size() int {
	return len(this.elements)
}


// https://groups.google.com/forum/#!topic/golang-nuts/UyKree3BCQ0
type StringQueue struct {
	Queue
}

func (s *StringQueue) Push(n string) {
	s.Queue.Push(n)
}
func (s *StringQueue) Pop() (string, error) {
	tmp, err := s.Queue.Pop()
	var val string
	if tmp != nil {
		val = tmp.(string)
	}
	return val, err
}

// https://groups.google.com/forum/#!topic/golang-nuts/UyKree3BCQ0
type IntQueue struct {
	Queue
}

func (s *IntQueue) Push(n int) {
	s.Queue.Push(n)
}
func (s *IntQueue) Pop() (int, error) {
	tmp, err := s.Queue.Pop()
	var val int
	if tmp != nil {
		val = tmp.(int)
	}
	return val, err
}