/**
leetcode 209 Minimum Size Subarray Sum

Given an array of positive integers nums and a positive integer target,
return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr]
of which the sum is greater than or equal to target. If there is no such subarray,
return 0 instead.

*/
package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	start := 0
	res := math.MaxInt32
	sum := 0
	for end := 0;end < len(nums);end++ {
		sum += nums[end]
		for sum >= target {
			res = minInt(res, end - start + 1)
			sum -= nums[start]
			start++
		}

	}
	if res == math.MaxInt32 {
		return 0
	} else {
		return res
	}

}

