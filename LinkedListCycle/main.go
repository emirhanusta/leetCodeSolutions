package main

import "fmt"

type ListNode struct {
	val  int
	Next *ListNode
}

func main() {
	listNode1 := &ListNode{
		val:  1,
		Next: nil,
	}
	fmt.Println(hasCycle(listNode1))
}

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]bool)
	currentNode := head
	for currentNode != nil {
		if nodes[currentNode] {
			return true
		}
		nodes[currentNode] = true
		currentNode = currentNode.Next
	}
	return false
}
