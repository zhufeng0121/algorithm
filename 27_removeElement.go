/**
 leetcode 27 remove Element

 Given an array nums and a value val, remove all instances
 of that value in-place and return the new length.

 Do not allocate extra space for another array, you must
 do this by modifying the input array in-place with O(1) extra memory.

 The order of elements can be changed. It doesn't matter what
 you leave beyond the new length.

 */
package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	head, tail := 0, 0
	for tail < len(nums) {
		if nums[tail] == val {
			tail++
		} else {
			nums[head] = nums[tail]
			head++
			tail++
		}
	}
	return head

}
