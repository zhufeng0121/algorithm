/**
 leetcode 162 Find Peak Element

 A peak element is an element that is strictly greater than its neighbors.

 Given an integer array nums, find a peak element, and return its index.
 If the array contains multiple peaks, return the index to any of the peaks.

 You may imagine that nums[-1] = nums[n] = -∞.

 You must write an algorithm that runs in O(log n) time.

 */
package main
// 此题用二分法 其实就是寻找极值的一个过程
//PS 记得研一上算法课时候，老师讲过 ，而且考试考过 ，不光有一维的，而且还有二维的数组
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	start, end := 0, len(nums) - 1
	for start < end {
		mid := start + (end - start) >> 1
		if nums[mid] > nums[mid +1] {
			end = mid
		} else {
			start = mid + 1
		}

	}
	return start

}
