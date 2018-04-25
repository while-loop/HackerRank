package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii
func main() {
	cases := []struct {
		list *ListNode
		exp  *ListNode
	}{
		{list(1, 1, 1, 2, 3), list(2, 3)},
		{list(1, 2, 3, 3, 4, 4, 5), list(1, 2, 5)},
	}

	for _, tc := range cases {
		act := deleteDuplicates(tc.list)
		if reflect.DeepEqual(tc.exp, act) {
			fmt.Printf("pass\n%v == %v\n", tc.exp, act)
		} else {
			fmt.Printf("fail:\ndeleteDuplicates(%v) = %v - want: %v\n", tc.list, act, tc.exp)
		}
		fmt.Println()
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	seen := map[int]struct{}{}
	root := head
	var prev *ListNode
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			seen[head.Val] = struct{}{}
			if prev == nil {
				prev = head.Next
				head = head.Next
			}

			prev.Next = head.Next
			head = prev
		} else if _, exists := seen[head.Val]; exists {
			prev.Next = head.Next
			head = prev
		}

		prev = head
		head = head.Next
	}

	return root
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
