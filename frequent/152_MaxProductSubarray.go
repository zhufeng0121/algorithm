/**
 leetcode 152 Maximum Product Subarray

 Given an integer array nums, find a contiguous non-empty subarray within
 the array that has the largest product, and return the product.

 It is guaranteed that the answer will fit in a 32-bit integer.

 A subarray is a contiguous subsequence of the array.

 */
package main

// 所有子数组 都会以某个元素中某个位置结束，如果求出以每一个位置结尾的子数组最大的累乘积
// 结果 max {以arr[0],arr[1] ... arr[n] 结尾的所有子数组的最大累乘积}
//  DP ：
// 假设 arr[i-1] 结尾的最大子数组累乘积为max,最小为min  那么以arr[i] 结尾的最大子数组累乘积
// 最大为 max(arr[i] * min,arr[i] * max,arr[i])

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	max, min := nums[0], nums[0]
	res := nums[0]
	maxEnd, minEnd := 0, 0;

	for i := 1; i < len(nums);i++ {
		maxEnd = nums[i] * max
		minEnd = nums[i] * min
		max = maxInt(maxInt(maxEnd,minEnd),nums[i])
		min = minInt(minInt(maxEnd,minEnd),nums[i])
		//max 代表以 nums[i]元素结尾的最大的子数组乘积的值，还需要 一起比较
		res = maxInt(res,max)

	}
	return res

}

func maxProductII() int {
	//TODO:
	return 0
}

func maxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}