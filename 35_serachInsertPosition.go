/**
 leetcode 35 Search Insert Position

 Given a sorted array of distinct integers and a target value, return the index
 if the target is found. If not, return the index where it would be if it were
 inserted in order.

 You must write an algorithm with O(log n) runtime complexity.
 */
package main

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	return start

}
