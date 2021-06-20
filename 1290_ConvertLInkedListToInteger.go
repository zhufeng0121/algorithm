/**
 leetcode 1290.Convert Binary Number in a Linked List to Integer

 Given head which is a reference node to a singly-linked list.
 The value of each node in the linked list is either 0 or 1.
 The linked list holds the binary representation of a number.

 Return the decimal value of the number in the linked list.

*/
package main

//秦九韶算法
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res << 1 + head.Val
		head = head.Next
	}
	return res

}
