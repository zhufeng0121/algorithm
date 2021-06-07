/**
 leetcode 704 Binary Search

 Given an array of integers nums which is sorted in ascending order,
 and an integer target, write a function to search target in nums.
 If target exists, then return its index. Otherwise, return -1.

 You must write an algorithm with O(log n) runtime complexity.

 */
package main

//二分法的题
func searchI(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}


