/**
 leetcode 206 Reverse Linked List

 Given the head of a singly linked list,
 reverse the list, and return the reversed list.

*/
package main

//Easy: 经典题 ： 链表反转
// 可以采用递归的方法，也可以采用头插法

//迭代法
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		//next 记录当前节点的Next节点
		next := curr.Next
		//prev赋值 curr.Next
		curr.Next = prev
		//将prev 赋值为 curr
		prev = curr
		// curr 赋值为 next
		curr = next
	}
	return prev
}
