package collection

import "errors"

/**
	Kinda generic Stack
 */
type Stack struct {
	elements []interface{}
}

func (this *Stack) Empty() bool {
	return len(this.elements) == 0
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
		size := len(this.elements) - 1
		res := this.elements[size]
		this.elements[size] = nil
		this.elements = this.elements[:size]
		return res, nil
	}
}

func (this *Stack) Push(val interface{}) {
	this.elements = append(this.elements, val)
}

func (this *Stack) Size() int {
	return len(this.elements)
}


// https://groups.google.com/forum/#!topic/golang-nuts/UyKree3BCQ0
type StringStack struct {
	Stack
}

func (s *StringStack) Push(n string) {
	s.Stack.Push(n)
}
func (s *StringStack) Pop() (string, error) {
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
func (s *IntStack) Pop() (int, error) {
	tmp, err := s.Stack.Pop()
	var val int
	if tmp != nil {
		val = tmp.(int)
	}
	return val, err
}