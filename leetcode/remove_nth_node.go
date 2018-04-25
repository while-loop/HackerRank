package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	cases := []struct {
		list *ListNode
		n    int
		exp  *ListNode
	}{
		{list(1, 2), 2, list(2)},
		{list(1, 2, 3, 4, 5), 2, list(1, 2, 3, 5)},
		{list(1), 1, nil},
	}

	for _, tc := range cases {
		act := removeNthFromEnd(tc.list, tc.n)
		if reflect.DeepEqual(tc.exp, act) {
			fmt.Printf("pass\n%v == %v\n", tc.exp, act)
		} else {
			fmt.Printf("fail\n%v != %v\n", tc.exp, act)
		}
		fmt.Println()
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur, ahead := head, head
	var prev *ListNode
	for i := 0; i < n; i++ {
		ahead = ahead.Next
	}

	for ahead != nil {
		prev = cur
		cur = cur.Next
		ahead = ahead.Next
	}

	if prev == nil {
		head = head.Next
	} else {
		prev.Next = cur.Next
	}

	return head
}

func list(nums ...int) *ListNode {
	tmp := &ListNode{}
	head := tmp
	for idx, v := range nums {
		if idx < len(nums)-1 {
			tmp.Next = &ListNode{}
		}

		tmp.Val = v
		tmp = tmp.Next
	}
	return head
}

func (ln ListNode) String() string {
	l := &ln

	s := fmt.Sprintf("[%d", l.Val)
	l = l.Next
	for l != nil {
		s += fmt.Sprintf(" %d", (*l).Val)
		l = l.Next
	}
	s += "]"
	return s
}
