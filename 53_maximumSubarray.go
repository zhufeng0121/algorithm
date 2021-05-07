/**
 leetcode 53 Maximum Subarray

 Given an integer array nums, find the contiguous subarray
 (containing at least one number) which has the largest sum
 and return its sum.
 */
package main

const INT_MAX = int(^uint32(0) >> 1)

func maxSubArray(nums []int) int {
	max := ^INT_MAX
	cur := 0
	for i := 0; i < len(nums);i++ {
		cur += nums[i]
		max = maxInt(cur,max)
		if cur < 0 {
			cur = 0
		}
	}
	return max

}

func maxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}



