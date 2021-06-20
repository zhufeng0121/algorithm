/**

 You are given the head of a singly linked-list. The list
 can be represented as:

 L0 → L1 → … → Ln - 1 → Ln
 Reorder the list to be on the following form:

 L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
 You may not modify the values in the list's nodes. Only nodes
themselves may be changed.

 */
package main

//找中点 + 反转后半部分 + 合并前后两部分
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//断开中点，反转后半部分
	



}
