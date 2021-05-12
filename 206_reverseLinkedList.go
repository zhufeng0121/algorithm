/**
 leetcode 206 Reverse Linked List

 Given the head of a singly linked list,
 reverse the list, and return the reversed list.

*/
package main

//Easy: 经典题 ： 链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, temp := head, head
	for cur.Next != nil {
		temp = cur.Next
		cur.Next = cur.Next.Next
		temp.Next = cur

	}



}
