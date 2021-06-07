/**
 leetcode 141 Linked List Cycle

 Given head, the head of a linked list, determine if the linked
 list has a cycle in it.

 There is a cycle in a linked list if there is some node in the list that
 can be reached again by continuously following the next pointer. Internally,
 pos is used to denote the index of the node that tail's next pointer is
 connected to. Note that pos is not passed as a parameter.

 Return true if there is a cycle in the linked list. Otherwise, return false.

 */
package main

//判断链表是否有环存在
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	// 这样写存在一个问题，就是 初始的时候，slow 和 fast 都在head ，那么此时slow 和 fast相等，直接返回true
	for slow != nil && fast != nil {
		if slow != fast {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return true
		}

	}
	return false

}


//判断链表是否有环存在(这个错了好多遍，不应该，要搞懂)
func hasCycleI(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	// 这样写存在一个问题，就是 初始的时候，slow 和 fast 都在head ，那么此时slow 和 fast相等，直接返回true
	for slow != fast {
		//要确保fast以及fast.next存在，这样fast指针才能走两步。 如果直接使用fast.next.next，
		//如果fast以及fast.next不存在，会直接报错。 又因为此处fast比slow快，所以可以不用加slow。
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
	return true

}

//可以使用类似于do-while 循环的方式，这样 slow 和fast 先走，然后再判断fast和slow是否相等

func hasCycleII(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

