/**
 leetcode 81 Search in Rotated Sorted Array II

 There is an integer array nums sorted in non-decreasing order
 (not necessarily with distinct values).

 Before being passed to your function, nums is rotated at an unknown
 pivot index k (0 <= k < nums.length) such that the resulting array
 is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
 (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot
 index 5 and become

 Given the array nums after the rotation and an integer target, return true
 if target is in nums, or false if it is not in nums.

 You must decrease the overall operation steps as much as possible.

 */
package main

// using Binary Search  [4,5,6,6,7,0,1].
func searchII(nums []int, target int) bool {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) >> 1
		if target == nums[mid] {
			return true
		}
		//根据 nums[mid] 和 nums[left]来进行判断mid是落在了左边的有序数组还是右边的有序数组
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else if nums[mid] < target || target < nums[left]  {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else if target < nums[mid] || target > nums[right] {
				right = mid -1

			}
		}
	}
	return false

}


