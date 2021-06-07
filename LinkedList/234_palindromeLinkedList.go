/**
 leetcode 234. Palindrome Linked List

 Given the head of a singly linked list, return true if it is a palindrome.

*/
package LinkedList

//判断一个链表是否是回文链表 ，此题方法也特别多，很明显。可以用栈
func isPalindromeI(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	array := make([]int,0)
	cur := head
	for cur != nil {
		array = append(array,cur.val)
		cur = cur.Next
	}
	for i := 0;i < len(array) >> 1;i++ {
		if array[i] != array[len(array) - i -1] {
			return false
		}
	}
	return true

}
//再考虑一下其他的做法
func isPalindromeII(head *ListNode) bool {
	//TODO:
	return false

}
