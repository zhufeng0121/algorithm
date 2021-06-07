/**
 leetcode 142 Linked List Cycle II

 Given a linked list, return the node where the cycle begins.
 If there is no cycle, return null.

 There is a cycle in a linked list if there is some node in the list
 that can be reached again by continuously following the next pointer.
 Internally, pos is used to denote the index of the node that tail's next
 pointer is connected to. Note that pos is not passed as a parameter.

 Notice that you should not modify the linked list.

 */
package LinkedList

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil

}

