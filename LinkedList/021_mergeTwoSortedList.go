/**
 leetcode 21 Merge Two Sorted Lists

 Merge two sorted linked lists and return it as a sorted list.
 The list should be made by splicing together the nodes of the first two lists.

 */
package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{val: 0,Next: nil}
	current := head
	for l1 != nil && l2 != nil {
		if l1.val < l2.val {
			current.Next = &ListNode{val: l1.val,Next: nil}
			l1 = l1.Next
			current = current.Next
		} else {
			current.Next = &ListNode{val: l2.val,Next: nil}
			l2 = l2.Next
			current = current.Next
		}

		if l1 == nil {
			for l2 != nil {
				current.Next = &ListNode{val: l2.val,Next: nil}
				l2 = l2.Next
				current = current.Next
			}
		}
		if l2 == nil {
			for l1 != nil {
				current.Next = &ListNode{val: l1.val,Next: nil}
				l1 = l1.Next
				current = current.Next
			}
		}
	}
	return head.Next
}

func mergeTwoListsII(l1 *ListNode, l2 *ListNode) *ListNode {
	//TODO: 可以用递归的写法再试一下
	return nil
}
