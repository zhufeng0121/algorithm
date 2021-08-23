/**
 leetcode 795 Number of Subarrays with Bounded Maximum

 Given an integer array nums and two integers left and right,
 return the number of contiguous non-empty subarrays such that
 the value of the maximum array element in that subarray is
 in the range [left, right].

 The test cases are generated so that the answer will fit in a
 32-bit integer.
 */

package main

//此题可以利用滑动窗口 + 前缀和 前缀和的知识可以见帖子
// https://blog.csdn.net/fhqfjevfhp/article/details/118215382

//该题的思路是，寻找不大于right子数组的个数，在寻找不大于left-1的子数组的个数，相减即可
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return countSubArray(nums,right) - countSubArray(nums,left-1)

}

func countSubArray(nums []int, k int) int {
	ans := 0
	pre := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= k {
			pre++
		}else {
			pre = 0
		}
		ans += pre

	}
	return ans
}