/**
 leetcode 876 Middle of the Linked List

 Given a non-empty, singly linked list with head node head,
 return a middle node of linked list.

 If there are two middle nodes, return the second middle node.

 */
package LinkedList

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}
