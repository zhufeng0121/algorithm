/**
 leetcode 268. Missing Number

 Given an array nums containing n distinct numbers in the range [0, n],
 return the only number in the range that is missing from the array.

 Follow up: Could you implement a solution using only O(1) extra space
 complexity and O(n) runtime complexity?

 */
package main

import "sort"

//n distinct numbers [0,n]共 n + 1 个数 [o,n] sum  n * (n + 1) / 2 n == len(nums)
func missingNumber(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	actualSum := 0
	for _, v := range nums {
		actualSum += v
	}
	return expectedSum - actualSum

}

//异或操作满足结合率
func missingNumberI(nums[] int) int {
	expect := len(nums)
	for i, v := range nums {
		expect = expect ^ i ^ v
	}
	return expect

}
//利用二分法
func missingNumberII(nums[] int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	start, end := 0, len(nums)
	mid := 0
	//二分的区间是一个左闭右开的区间，因此是小于
	for start < end {
		mid = start + (end - start) >> 1
		if nums[mid] > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func missingNumberIII(nums[] int) int {
	res := len(nums)
	for i, v := range nums {
		res += i - v
	}
	return res
}

