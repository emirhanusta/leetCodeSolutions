package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := mergeTwoLists(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	if list1.Val < list2.Val {
		newList = list1
		list1 = list1.Next
	} else {
		newList = list2
		list2 = list2.Next
	}
	current := newList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	}

	if list2 == nil {
		current.Next = list1
	}

	return newList
}
