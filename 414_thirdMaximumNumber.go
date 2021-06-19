/**
 leetcode 414 Third Maximum Number

 Given integer array nums, return the third maximum number in this array.
 If the third maximum does not exist, return the maximum number.

 */
package main

import "math"

func thirdMax(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	// arr[0] > arr[1] > arr[2]
	//如果嫌建堆麻烦，可以利用数组保存前3个最大的数
	arr := make([]int,3)

	for i := 0; i < 3;i++ {
		arr[i] = math.MinInt64
	}
	for i := 0; i < len(nums);i++ {
		if nums[i] > arr[0] {
			arr[2] = arr[1]
			arr[1] = arr[0]
			arr[0] = nums[i]
		} else if nums[i] > arr[1] && nums[i] < arr[0] {
			arr[2] = arr[1]
			arr[1] = nums[i]
		} else if nums[i] > arr[2] && nums[i] < arr[1] {
			arr[2] = nums[i]
		}
	}
	if arr[2] > math.MinInt64 {
		return arr[2]
	}
	//如果没有第三大的数，返回最大的数
	return arr[0]
}
