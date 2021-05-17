/**
 leetcode 33 Search in Rotated Sorted Array

 There is an integer array nums sorted in ascending order (with distinct values).

 Prior to being passed to your function, nums is rotated at an
 unknown pivot index k (0 <= k < nums.length) such that the resulting
 array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]
 (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and
 become [4,5,6,7,0,1,2].

 Given the array nums after the rotation and an integer target, return the index of
 target if it is in nums, or -1 if it is not in nums.

 You must write an algorithm with O(log n) runtime complexity.

 */
package main

//好题。说来不信，这个题居然是在虎扑上的算法号上看懂的
// 其实数组旋转后，就相当于原来的一个有序数组变成了两个有序数组 可以 通过 nums[left] 和 nums[mid]的
//大小，来进行比较 如果 nums[left] < nums[mid] 说明 mid 和 left在同一个有序数组中，如果 nums[left] > nums[mid]
// 说明nums[mid]在 另一个有序数组中

// 如果 mid 和 left在同一个数组中，那么 target 只能有两种可能性，1. nums[left] <= target < nums[mid]（之前已经判断过target和
//nums[mid]）是否相等 2. nums[mid] < target || target < target[left] .
// 如果是 1 ，则直接让right = mid -1 如果是2 让 left = mid + 1

//如果 mid 和 left 不在同一个数组中，target也存在两种可能性 1. target > nums[mid] && target <= nums[right]
// 2. target < nums[mid] || target >= nums[left]（这里应该携程 target > nums[right]）
// 如果是 1, left = mid + 1 如果是2 right = mid - 1
func search(nums []int, target int) int {
	left  := 0
	right := len(nums) - 1
	mid   := 0
	for left <= right {
		mid = left + (right - left) >> 1
		if nums[mid] == target {
			return mid
		}
		// mid 和 left元素处于同一个升序数组中
		if nums[mid] >= nums[left] {
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else if target > nums[mid] || target < nums[left] {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else if nums[mid] > target || target > nums[right] {
				right = mid - 1
			}

		}

	}
	return -1
}
