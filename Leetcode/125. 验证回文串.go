package main

import "fmt"

func main() {
	node := ListNode{
		Val:1,
	}
	node2 := ListNode{
		Val:2,
		Next:&node,
	}
	node3 := ListNode{
		Val:3,
		Next:&node2,
	}
	node4 := ListNode{
		Val:4,
		Next:&node3,
	}
	node5 := ListNode{
		Val:5,
		Next:&node4,
	}
	fmt.Println(reverseList(&node5))
}

/**
* Definition for singly-linked list.
*/
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}