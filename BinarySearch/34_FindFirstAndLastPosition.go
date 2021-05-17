/**
 leetcode 34 Find First and Last Position of Element In Sorted Array

 Given an array of integers nums sorted in ascending order, find the starting
 and ending position of a given target value.

 If target is not found in the array, return [-1, -1].

 You must write an algorithm with O(log n) runtime complexity.

*/
package main

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	if target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}
	start, end := 0, 0

	for left < right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			for i:= mid - 1; i >= 0;i-- {
				if nums[i] != target {
					if i == 0 {
						start = i
					} else {
						start = i + 1
					}
					break
				}
			}
			for i := mid + 1;i <= right;i++ {
				if nums[i] != target {
					if i == right {
						end = i
					} else {
						end = i - 1
					}
					break
				}
			}
			return []int{start, end}

		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return []int{-1, -1}

}

// 运行两次二分法，一次找最左，一次找最右
func searchRangeII(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	if target < nums[left] || target > nums[right] {
		return []int{-1, -1}
	}
	left = searchFirstEqualElement(nums,target)
	right = searchLastEqualElement(nums,target)
	return []int{left,right}

}
// 二分查找第一个与 target 相等的元素,时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		}

	}
	return -1

}
// 二分查找最后一个与 target 相等的元素,时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			low = mid + 1
		}
	}
	return -1

}

//还有一种解法
func searchRangeIII() int{
	//TODO:
	return 0
}


