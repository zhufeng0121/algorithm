/**
 leetcode 26. remove Duplicates from Sorted Array

 Given a sorted array nums, remove the duplicates
 in-place such that each element appears only once
 and returns the new length.

 Do not allocate extra space for another array, you
 must do this by modifying the input array in-place
 with O(1) extra memory.
 */
package main

/**
 思路 : 用两个指针，一个指针不断向前走，若遇到不同元素，交换，相同，继续加1
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	end := 0
	for end < len(nums) {
		if nums[end] == nums[start] {
			end++
		} else {
			start++
			nums[start] = nums[end]
			end++
		}
	}
	return start + 1

}
