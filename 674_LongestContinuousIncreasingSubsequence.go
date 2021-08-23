/**
 leetcode 674 Longest Continuous Increasing Subsequence

 Given an unsorted array of integers nums, return the length of
 the longest continuous increasing subsequence (i.e. subarray).
 The subsequence must be strictly increasing.

 A continuous increasing subsequence is defined by two indices
 l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1],
 nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].
 */
package main
//获取最长递增连续子序列
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res++
		} else {
			res = 1
		}
		if res > max {
			max = res
		}
	}
	return max

}
