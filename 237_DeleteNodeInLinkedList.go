/**
 leetcode 237 Delete Node in a Linked List

 Write a function to delete a node in a singly-linked list.
 You will not be given access to the head of the list, instead
 you will be given access to the node to be deleted directly.

 It is guaranteed that the node to be deleted is not a tail node
 in the list.

 */
package main

func deleteNode(node *ListNode) {
	value := node.Next.val
	node.Next = node.Next.Next
	node.val = value

}
