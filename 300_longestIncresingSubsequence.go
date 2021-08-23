/**
 leetcode 300 Longest Increasing Subsequence

 Given an integer array nums, return the length of the longest
 strictly increasing subsequence.

 A subsequence is a sequence that can be derived from an array by
 deleting some or no elements without changing the order of the
 remaining elements. For example, [3,6,2,7] is a subsequence of
 the array [0,3,1,6,2,2,7].
 */
package main

//最长递增子序列
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0;j < i;j++ {
			if nums[i] > nums[j] && dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return maxInSlice(dp)
}
func maxInSlice(dp []int) int {
	for i := 1;i < len(dp);i++ {
		if dp[i] > dp[0] {
			dp[i] , dp[0] = dp[0], dp[i]
		}
	}
	return dp[0]
}