/**
 leetcode 153 Find Minimum in Rotated Sorted Array

 Suppose an array of length n sorted in ascending order is rotated between
 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

 [4,5,6,7,0,1,2] if it was rotated 4 times.
 [0,1,2,4,5,6,7] if it was rotated 7 times.
 Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

 Given the sorted rotated array nums of unique elements, return the minimum element of this array.

 You must write an algorithm that runs in O(log n) time.


 */
package main
//非常非常经典的一道题，binarySearch的应用，一定要好好理解理解
func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	start, end := 0, len(nums)-1
	mid := 0
	for start < end {
		mid = start + (end - start) >> 1
		//如果中间数小于最右边的数，那么说明最小数一定在左半边，至于mid 是不是最小，需要end = mid 保留
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return nums[start]

}
