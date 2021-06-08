/**
 leetcode 203 Remove Linked List Elements

 Given the head of a linked list and an integer val, remove all the nodes
 of the linked list that has Node.val == val, and return the new head.

 */
package main

import "go/ast"

//可改进
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	virtnode := &ListNode{Next: head}
	before := virtnode
	cur := virtnode.Next
	for cur != nil {
		if cur.val == val {
			before.Next = cur.Next
		} else {
			before = before.Next
		}
		cur = cur.Next
	}
	return virtnode.Next

}

//递归的做法
func removeElementsII(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next,val)
	if head.val == val {
		return head.Next
	}
	return head
}

func removeElementsIV(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	
}

