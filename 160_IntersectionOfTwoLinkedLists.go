/**
 leetcode 160 Intersection of Two Linked Lists

 Given the heads of two singly linked-lists headA and headB,
 return the node at which the two lists intersect. If the two
 linked lists have no intersection at all, return null.

 */
package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	cur := headA
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = headA
	slow, fast := headB, headB
	for fast.Next != nil {
		for slow != fast {
			slow = slow.Next
			fast = fast.Next.Next
		}
		//相遇
		temp := slow
		fast = headB
		for slow != fast {
			fast = fast.Next
			slow = slow.Next
		}
		for temp.Next != headA {
			temp = temp.Next
		}
		temp.Next = nil
		return fast
	}
	return nil

}
// 判断一个链表是否
func hasCycle(head *ListNode) bool {

}