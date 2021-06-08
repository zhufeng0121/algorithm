/**
 leetcode 23 Merge k Sorted Lists

 You are given an array of k linked-lists lists, each linked-list
 is sorted in ascending order.

 Merge all the linked-lists into one sorted linked-list and return it.

 */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Brute-Force
//Traverse all the linked lists and collect the values of the nodes into an array.
//Sort and iterate over this array to get the proper value of nodes.
//Create a new sorted linked list and extend it with the new nodes.
func mergeKLists(lists []*ListNode) *ListNode {
	return nil

}
