package collection

import "errors"

/**
	Kinda generic Stack
 */
type Stack struct {
	elements []interface{}
	size     int
}

func (this *Stack) Empty() bool {
	return this.size == 0
}

func (this *Stack) Peek() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Stack")
	} else {
		return this.elements[0], nil
	}
}

func (this *Stack) Pop() (interface{}, error) {
	if this.Empty() {
		return nil, errors.New("Empty Stack")
	} else {
		this.size-- // decrement before return because of 0-index
		ele := this.elements[this.size]
		this.elements[this.size] = nil
		return ele, nil
	}
}

func (this *Stack) Push(val interface{}) {
	this.elements = append(this.elements, val)
	this.size++
}

func (this *Stack) Size() int {
	return this.size
}


// https://groups.google.com/forum/#!topic/golang-nuts/UyKree3BCQ0
type StringStack struct {
	Stack
}

func (s *StringStack) Push(n string) {
	s.Stack.Push(n)
}
func (s *StringStack) Pop() (string, error){
	tmp, err := s.Stack.Pop()
	var val string
	if tmp != nil {
		val = tmp.(string)
	}
	return val, err
}

// https://groups.google.com/forum/#!topic/golang-nuts/UyKree3BCQ0
type IntStack struct {
	Stack
}

func (s *IntStack) Push(n int) {
	s.Stack.Push(n)
}
func (s *IntStack) Pop() (int, error){
	tmp, err := s.Stack.Pop()
	var val int
	if tmp != nil {
		val = tmp.(int)
	}
	return val, err
}