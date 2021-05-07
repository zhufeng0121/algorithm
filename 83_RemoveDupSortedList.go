/**
 leetcode 83 Remove Duplicates from Sorted List
 */

package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

func deleteDuplicates(head *ListNode2) *ListNode2 {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}

	return head
}
