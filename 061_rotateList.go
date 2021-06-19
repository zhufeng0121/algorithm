/**
 leetcode 61 Rotate List

 Given the head of a linked list, rotate the list to the right by k places.

 */
package main
//旋转k（k < len）就是将尾部k个节点 单独截取出来，分成两个链表，然后将尾部的子链表和头部的子链表连接起来
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	length := 1
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	tail := cur
	k = k % length
	if k == 0 {
		return head
	}
	newTail := head
	count := length - k - 1
	for count != 0 {
		newTail = newTail.Next
		count--
	}
	newHead := newTail.Next
	tail.Next = head
	newTail.Next = nil
	return newHead

}