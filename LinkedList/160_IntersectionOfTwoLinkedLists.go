/**
 leetcode 160 Intersection of Two Linked Lists

 Given the heads of two singly linked-lists headA and headB,
 return the node at which the two lists intersect. If the two
 linked lists have no intersection at all, return null.

 */
package LinkedList

type ListNode struct {
	val int
	Next *ListNode
}

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
/**
 要求时间复杂度为 O(N)，空间复杂度为 O(1)。如果不存在交点则返回 null。

 设 A 的长度为 a + c，B 的长度为 b + c，其中 c 为尾部公共部分长度，可知 a + c + b = b + c + a。

 当访问 A 链表的指针访问到链表尾部时，令它从链表 B 的头部开始访问链表 B；同样地，
 当访问 B 链表的指针访问到链表尾部时，令它从链表 A 的头部开始访问链表 A。这样就能
 控制访问 A 和 B 两个链表的指针能同时访问到交点。

 如果不存在交点，那么 a + b = b + a，以下实现代码中 l1 和 l2 会同时为 null，从而退出循环。
 */
func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a

}

